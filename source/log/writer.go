package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

func timeDivisionWriter(path string) io.Writer {
	return &lumberjack.Logger{
		Filename:   path + ".log", //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    2,             //文件大小限制,单位MB
		MaxBackups: 100,           //最大保留日志文件数量
		MaxAge:     7,             //日志文件保留天数
		Compress:   true,          //是否压缩处理
	}
}
