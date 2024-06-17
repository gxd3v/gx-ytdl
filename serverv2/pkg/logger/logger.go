package logger

import (
	"context"
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func Init(_ context.Context) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Logger = Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &Logger
	Info().Msg("logger initialized successfully")
}

func Info() *zerolog.Event {
	return Logger.Info()
}

func Err(err error) *zerolog.Event {
	return Logger.Err(err)
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}

func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

func Fields(f map[string]any) zerolog.Logger {
	return Logger.With().Fields(f).Logger()
}
