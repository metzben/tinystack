package api

import (
	"github.com/metzben/tinystack/internal/api/url"
	"net/http"
)

func (app *Application) BuildRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(url.Home, app.Home)
	mux.HandleFunc(url.Users, app.HandleUserName)
	mux.HandleFunc(url.Messages, app.anthropicMessages)

	return mux
}
