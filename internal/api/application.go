package api

import (
	"fmt"
	"github.com/metzben/tinystack/internal/config"
	"github.com/rs/zerolog"
	"net/http"
)

type Application struct {
	Logger        zerolog.Logger
	Configuration config.Configuration
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msgf("status code %v", r.Response.StatusCode)

	fmt.Fprintln(w, "yo we have a go app running!")
}

func (app *Application) HandleUserName(w http.ResponseWriter, r *http.Request) {
	// HOW LOG?
	// HOW Configuration??
	name := r.PathValue("name")
	fmt.Fprintln(w, "yo name is: ", name)
}
