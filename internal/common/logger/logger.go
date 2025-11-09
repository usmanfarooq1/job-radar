package logger

import (
	"github.com/rs/zerolog"
)

func InitLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return *zerolog.DefaultContextLogger
}
