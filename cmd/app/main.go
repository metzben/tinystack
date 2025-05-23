package main

import (
	"os"

	"github.com/metzben/tinystack/internal/api"
	"github.com/metzben/tinystack/internal/config"
	"github.com/metzben/tinystack/internal/secrets"
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

	configuration, loadErr := config.Load(envFile)
	if loadErr != nil {
		log.Fatal().Err(loadErr).Msg("could not read .env file")
	}

	secretMgr, secretMgrErr := secrets.NewGoogleSecretsClient(configuration.GCPProjectID)
	if secretMgrErr != nil {
		log.Fatal().Err(secretMgrErr).Msg("cannot load secret")
	}

	app := api.Application{
		Logger:        log,
		Configuration: configuration,
		SecretManager: secretMgr,
	}

	// start the server
	serverErr := app.Serve()

	if serverErr != nil {
		log.Fatal().Msgf("server bonk err: %s", serverErr)
	}
}
