package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZeroErrorLog(dirName string) *zerolog.Logger {
	writer := timeDivisionWriter(dirName + "/error")
	consoleWriter := zerolog.ConsoleWriter{
		Out:     writer,
		NoColor: true,
	}
	logger := log.Output(consoleWriter).With().Logger()
	return &logger
}

func NewZeroAccessLog(dirName string) *zerolog.Logger {
	writer := timeDivisionWriter(dirName + "/access")
	consoleWriter := zerolog.ConsoleWriter{
		Out:             writer,
		NoColor:         true,
		FormatTimestamp: func(i interface{}) string { return "" },
		FormatLevel:     func(i interface{}) string { return "" },
	}
	logger := log.Output(consoleWriter).With().Logger()
	return &logger
}
