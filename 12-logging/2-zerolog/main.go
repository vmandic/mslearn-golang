package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msg("Hey, I'm a log message 1!")
	log.Debug().Msg("Hey, I'm a log message 2!")
}
