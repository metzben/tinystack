package main

import (
	"fmt"
	"net/http"
	"os"

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

	// new http server
	mux := http.NewServeMux()

	// setup routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/users", handleUsers)
	mux.HandleFunc("/users/{name}", handleUserName)

	log.Info().Msgf("starting server on port: %s", os.Getenv("PORT"))

	// start the server
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		log.Fatal().Msgf("server bonk err: %s", err)
	}
}

func handleUserName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintln(w, "yo name is: ", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yo we have a go app running!")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yo we have users")
}
