package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var dictLogLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

var (
	ZapErrorLogger  = NewZapErrorLogger("info")
	ZapAccessLogger = NewZapAccessLogger()
)

func NewZapErrorLogger(level string) *zap.Logger {
	zapLevel, ok := dictLogLevel[level]
	if !ok {
		zapLevel = zapcore.InfoLevel
	}
	writer := getLumberjackWriter(ErrorPath)
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()
	encoderConfig.CallerKey = "caller"
	encoderConfig.TimeKey = "ts"
	encoderConfig.LevelKey = "level"
	encoder := zapcore.NewConsoleEncoder(encoderConfig)     //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	return zap.New(core, zap.AddCaller())
}

func NewZapAccessLogger() *zap.Logger {
	writer := getLumberjackWriter(AccessPath)
	sink := zapcore.AddSync(writer)
	writeSyncer := zapcore.NewMultiWriteSyncer(sink)
	encoderConfig := GetEncoderConfig()                               //指定时间格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	return zap.New(core)
}

func GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		NameKey:        "logger",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(TimeFormat),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
