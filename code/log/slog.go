package log

import (
	"log/slog"
	"os"
	"runtime"
	"strconv"
)

var (
	accessLog = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	errorLog  = slog.New(slog.NewTextHandler(os.Stderr, nil))
)

func Access(msg string, args ...any) {
	accessLog.Info(msg, args...)
}

func Error(msg string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	errorLog.With(slog.String("call", file+":"+strconv.FormatInt(int64(line), 10))).Error(msg, args...)
}
