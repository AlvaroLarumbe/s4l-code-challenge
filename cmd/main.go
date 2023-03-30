package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"s4l-code-challenge/infrastructure/api"
	"time"
)

func main() {
	configureLogger()

	log.Info().Msg("API started!")

	httpServer := api.NewServer("0.0.0.0:8090")
	err := httpServer.Start()

	if err != nil {
		log.Error().Err(err).Msg("Cannot start HTTP server")

		return
	}
}

func configureLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
