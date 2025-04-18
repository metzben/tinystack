package api

import (
	"fmt"
	"net/http"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msg("home route hit")
	fmt.Fprintln(w, "yo we have a go app running!")
}

func (app *Application) HandleUserName(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msg("user route hit")
	name := r.PathValue("name")
	fmt.Fprintln(w, "yo name is: ", name)
}

// api endpoint that will accept a user prompt
func (app *Application) anthropicMessages(w http.ResponseWriter, r *http.Request) {
	// need to design golang struct for this request
	// unmarshal json to struct
	// we need to call the anthropic api
	// we need marshal the response from anthropic to a golang struct
	// need to reply with json
}
