package source

import (
	_ "embed"
	"go-server-gen/writer"
)

var (
	//go:embed log/writer.go
	LoggerCode string

	//go:embed log/zero.go
	ZeroLoggerCode string

	//go:embed log/slog.go
	SlogCode string

	//go:embed gin/bind.go
	GinBindCode string

	//go:embed hertz/bind.go
	HertzBindCode string
)

func GetDefaultCode() map[string]writer.WriteCode {
	const writeType = "overwrite"
	return map[string]writer.WriteCode{
		"pkg/log/writer.go": {
			File:  "pkg/log/writer.go",
			Write: writeType,
			Code:  LoggerCode,
		},
		"pkg/log/zero.go": {
			File:  "pkg/log/zero.go",
			Write: writeType,
			Code:  ZeroLoggerCode,
		},
		"biz/controller/internal/gin_bind.go": {
			File:  "biz/controller/internal/gin_bind.go",
			Write: writeType,
			Code:  GinBindCode,
		},
	}
}

func ModifyPackageName(src string, name string) string {
	return src
}
