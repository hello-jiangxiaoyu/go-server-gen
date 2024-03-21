package source

import (
	"embed"
	"go-server-gen/conf"
	"go-server-gen/source/internal"
	"go-server-gen/template"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"strings"
)

type GlobalData struct {
	ProjectName string
	Pkg         map[string]string
}

// GenPackageCode 生成默认代码
func GenPackageCode(layout conf.LayoutConfig, prefix string, overwrite bool) (map[string]writer.WriteCode, error) {
	writeType := writer.Overwrite
	if !overwrite {
		writeType = writer.Skip
	}
	projectName, err := utils.GetProjectName()
	if err != nil {
		return nil, err
	}

	server := layout.ServerType
	initCode := make(map[string]writer.WriteCode)
	tplMap := map[string]string{
		prefix + "main.go":                         internal.MainCodeMap[server],
		prefix + "README.md":                       internal.Readme,
		prefix + "Dockerfile":                      internal.DockerFile,
		prefix + "Makefile":                        internal.Makefile,
		prefix + ".gitignore":                      internal.GitIgnore,
		prefix + "biz/register.go":                 internal.RegisterCode,
		prefix + "pkg/response/error_request.go":   internal.ErrorRequest,
		prefix + "pkg/response/error_sql.go":       internal.ErrorSql,
		prefix + "pkg/response/error_unknown.go":   internal.ErrorUnknown,
		prefix + "pkg/response/success.go":         internal.Success,
		prefix + "pkg/response/service_code.go":    internal.ServiceCode,
		prefix + "pkg/middleware/recover.go":       internal.MiddlewareMap["recover"],
		prefix + "pkg/middleware/request_id.go":    internal.MiddlewareMap["request"],
		prefix + "biz/controller/internal/bind.go": replacePackage(GetEmbedContent("internal/"+server+"/bind.go"), "internal"),
		prefix + "pkg/response/response.go":        replacePackage(GetEmbedContent("internal/"+server+"/response.go"), "response"),
		prefix + "pkg/log/logger.go":               GetEmbedContent("internal/log/" + layout.LogType + ".go"),
		prefix + "pkg/log/writer.go":               GetEmbedContent("internal/log/writer.go"),
		prefix + "pkg/orm/gorm.go":                 GetEmbedContent("internal/orm/gorm.go"),
		prefix + "pkg/orm/gorm_gen.go":             GetEmbedContent("internal/orm/gorm_gen.go"),
		prefix + "pkg/utils/jwks.go":               GetEmbedContent("internal/utils/jwks.go"),
		prefix + "pkg/utils/map_int64.go":          GetEmbedContent("internal/utils/map_int64.go"),
		prefix + "pkg/utils/strings.go":            GetEmbedContent("internal/utils/strings.go"),
		prefix + "pkg/utils/tools.go":              GetEmbedContent("internal/utils/tools.go"),
		prefix + "pkg/utils/uuid.go":               GetEmbedContent("internal/utils/uuid.go"),
	}
	for fileName, body := range tplMap {
		fileName, body, err = template.ParseSource(fileName, body, GlobalData{ProjectName: projectName, Pkg: layout.Pkg})
		if err != nil {
			return nil, err
		}

		if body != "" {
			initCode[fileName] = writer.WriteCode{
				File:  fileName,
				Code:  body,
				Write: writeType,
			}
		}
	}
	return initCode, nil
}

var (
	//go:embed internal
	code embed.FS
)

// GetEmbedContent 从embed中读取文件内容
func GetEmbedContent(path string) string {
	res, err := code.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(res)
}

// 替换go文件包名
func replacePackage(src string, pkgName string) string {
	lines := strings.SplitN(src, "\n", -1)
	lines[0] = "package " + pkgName
	return strings.Join(lines, "\n")
}
