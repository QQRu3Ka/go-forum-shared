package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var logger *zerolog.Logger

func InitLogger() {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger = &l
}

func GetLogger() *zerolog.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}
