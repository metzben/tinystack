package main

import (
	"net/http"
	"os"

	"github.com/metzben/tinystack/internal/api"
	"github.com/metzben/tinystack/internal/config"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Info().Msg("Hey we have logging working!")

	envFile, openErr := os.Open(".env")
	if openErr != nil {
		log.Fatal().Err(openErr).Msg("failed to load .env file")
	}

	loadErr := config.Load(envFile)
	if loadErr != nil {
		log.Fatal().Err(loadErr).Msg("could not read .env file")
	}

	configuration := config.BuildConfiguration()

	app := api.Application{
		Logger:        log,
		Configuration: configuration,
	}

	// new http server
	mux := http.NewServeMux()

	// setup routes
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/users/{name}", app.HandleUserName)

	log.Info().Msgf("starting server on port: %s", configuration.Port)

	// start the server
	err := http.ListenAndServe(":"+configuration.Port, mux)
	if err != nil {
		log.Fatal().Msgf("server bonk err: %s", err)
	}
}
