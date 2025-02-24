package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/metzben/tinystack/internal/api/url"
	"github.com/metzben/tinystack/internal/config"
	"github.com/rs/zerolog"
)

type Application struct {
	Logger        zerolog.Logger
	Configuration config.Configuration
	sync.WaitGroup
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msgf("status code %v", r.Response.StatusCode)

	fmt.Fprintln(w, "yo we have a go app running!")
}

func (app *Application) HandleUserName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintln(w, "yo name is: ", name)
}

func (app *Application) Serve() error {

	mux := http.NewServeMux()

	srvr := &http.Server{
		Addr:         app.Configuration.Port,
		IdleTimeout:  time.Second * 2,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		Handler:      app.BuildRoutes(mux),
	}

	shutDownErrChan := make(chan error)

	// will immediately be dispatched to the go scheduler
	go func() {
		runtime.Gosched()

		// create the shutdown channel and provide it to the Notify func
		// so that it can listen for provided signals
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

		// then we BLOCK... by attempting to read from the shutdown channel
		// this go routine stops right here until it gets some kind of shutdown
		// signal
		s := <-shutdown

		// if shutdown is initiated, then we run the following code
		app.Logger.Info().Msgf("shutting server down with os signal: %s", s.String())

		// we set the amount of time we are going to give for graceful shutdown here
		// should be set in .env eventually
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srvr.Shutdown(ctx)
		if err != nil {
			shutDownErrChan <- err
		}
		// RIGHT HERE IS WHERE WE SHUT STUFF DOWN
		// example: app.UserRepo.Db.Close()
		app.Logger.Info().Msg("shutting stuff DOWN!")

		// wait for background tasks to complete
		app.Wait()
		shutDownErrChan <- nil
	}()

	app.Logger.Info().Msgf("starting prod server on port: %+v", app.Configuration.Port)
	if prodServerErr := srvr.ListenAndServe(); prodServerErr != nil && prodServerErr != http.ErrServerClosed {
		return prodServerErr
	}

	// this is again a blocking read from the error channel
	shuttingErr := <-shutDownErrChan
	if shuttingErr != nil {
		return shuttingErr
	}

	app.Logger.Info().Msgf("server has stopped gracefully on port: %+v\n", app.Configuration.Port)
	return nil
}

func (app *Application) BuildRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(url.Home, app.Home)
	mux.HandleFunc(url.Users, app.HandleUserName)

	return mux
}
