package log

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	ZeroErrorLogger  = NewZeroErrorLog()
	ZeroAccessLogger = NewZeroAccessLog()
)

func NewZeroErrorLog() *zerolog.Logger {
	writer := getLumberjackWriter(ErrorPath)
	consoleWriter := zerolog.ConsoleWriter{
		Out:             writer,
		NoColor:         true,
		FormatTimestamp: func(i interface{}) string { return "" },
		FormatLevel:     func(i interface{}) string { return "" },
	}
	logger := log.Output(consoleWriter).With().Caller().Logger()
	return &logger
}

func NewZeroAccessLog() *zerolog.Logger {
	writer := getLumberjackWriter(AccessPath)
	logger := zerolog.New(writer).With().Logger()
	return &logger
}

func Info() *zerolog.Event {
	return ZeroErrorLogger.Info().Str("ts", time.Now().Format(TimeFormat))
}

func Warn() *zerolog.Event {
	return ZeroErrorLogger.Warn().Str("ts", time.Now().Format(TimeFormat))
}

func Error() *zerolog.Event {
	return ZeroErrorLogger.Error().Str("ts", time.Now().Format(TimeFormat))
}
