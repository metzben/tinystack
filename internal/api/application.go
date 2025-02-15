package api

import (
	"fmt"
	"github.com/metzben/tinystack/internal/api/url"
	"github.com/metzben/tinystack/internal/config"
	"github.com/rs/zerolog"
	"net/http"
	"sync"
)

type Application struct {
	Logger        zerolog.Logger
	Configuration config.Configuration
	WG            sync.WaitGroup
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msgf("status code %v", r.Response.StatusCode)

	fmt.Fprintln(w, "yo we have a go app running!")
}

func (app *Application) HandleUserName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintln(w, "yo name is: ", name)
}

func (app *Application) BuildRoutes(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc(url.Home, app.Home)
	mux.HandleFunc(url.Users, app.HandleUserName)

	return mux
}
