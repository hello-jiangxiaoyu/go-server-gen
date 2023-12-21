package code

import _ "embed"

var (
	//go:embed log/writer.go
	LoggerCode []byte

	//go:embed log/zero.go
	ZeroLoggerCode []byte

	//go:embed log/slog.go
	SlogCode []byte

	//go:embed gin/bind.go
	GinBindCode []byte

	//go:embed hertz/bind.go
	HertzBindCode []byte
)
