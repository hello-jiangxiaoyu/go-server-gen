package source

import (
	"embed"
	"go-server-gen/writer"
)

var (
	//go:embed internal
	code embed.FS
	//go:embed response.yaml
	responseTpl []byte
)

func GetDefaultCode() map[string]writer.WriteCode {
	const writeType = "overwrite"
	return map[string]writer.WriteCode{}
}

func ModifyPackageName(src string, name string) string {
	return src
}
