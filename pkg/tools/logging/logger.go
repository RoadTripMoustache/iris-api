// Package logging contains all the methods related to the logging.
package logging

import (
	"github.com/rs/zerolog"
	"os"
)

var (
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
)

func Error(err error, args map[string]interface{}) {
	log(logger.Error().Err(err), "", args)
}

func Fatal(err error, args map[string]interface{}) {
	log(logger.Fatal().Err(err), "", args)
}

func Info(msg string, args map[string]interface{}) {
	log(logger.Info(), msg, args)
}

func Warn(msg string, args map[string]interface{}) {
	log(logger.Warn(), msg, args)
}

func log(logEvent *zerolog.Event, msg string, args map[string]interface{}) {
	for key, value := range args {
		logEvent = logEvent.Any(key, value)
	}

	logEvent.Msg(msg)
}
