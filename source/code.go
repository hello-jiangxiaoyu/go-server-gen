package source

import (
	"embed"
	"go-server-gen/conf"
	"go-server-gen/source/internal"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"strings"
)

type GlobalData struct {
	ProjectName string
	Pkg         map[string]conf.Package
}

// GenPackageCode 生成默认代码
func GenPackageCode(layout conf.LayoutConfig, server string, log string, code map[string]writer.WriteCode) error {
	projectName, err := utils.GetProjectName()
	if err != nil {
		return err
	}
	pkg := layout.Pkg

	tplMap := map[string]string{
		"main.go":                         internal.MainCodeMap[server],
		"README.md":                       internal.Readme,
		"biz/register.go":                 internal.RegisterCode,
		"pkg/response/error_request.go":   internal.ErrorRequest,
		"pkg/response/error_sql.go":       internal.ErrorSql,
		"pkg/response/error_unknown.go":   internal.ErrorUnknown,
		"pkg/response/success.go":         internal.Success,
		"pkg/response/service_code.go":    internal.ServiceCode,
		"pkg/middleware/recover.go":       internal.MiddlewareMap["recover"],
		"biz/controller/internal/bind.go": replacePackage(GetEmbedContent("internal/"+server+"/bind.go"), "internal"),
		"pkg/response/response.go":        replacePackage(GetEmbedContent("internal/"+server+"/response.go"), "response"),
		"pkg/log/logger.go":               GetEmbedContent("internal/log/" + log + ".go"),
		"pkg/log/writer.go":               GetEmbedContent("internal/log/writer.go"),
		"pkg/orm/gorm.go":                 GetEmbedContent("internal/orm/gorm.go"),
		"pkg/orm/gorm_gen.go":             GetEmbedContent("internal/orm/gorm_gen.go"),
		"pkg/utils/jwks.go":               GetEmbedContent("internal/utils/jwks.go"),
		"pkg/utils/map_int64.go":          GetEmbedContent("internal/utils/map_int64.go"),
		"pkg/utils/strings.go":            GetEmbedContent("internal/utils/strings.go"),
		"pkg/utils/tools.go":              GetEmbedContent("internal/utils/tools.go"),
	}
	for fileName, body := range tplMap {
		fileName, body, err = utils.ParseSource(fileName, body, GlobalData{ProjectName: projectName, Pkg: pkg})
		if err != nil {
			return err
		}

		if body != "" {
			code[fileName] = writer.WriteCode{
				File:  fileName,
				Code:  body,
				Write: "skip",
			}
		}
	}
	return nil
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

func replacePackage(src string, pkgName string) string {
	lines := strings.SplitN(src, "\n", -1)
	lines[0] = "package " + pkgName

	return strings.Join(lines, "\n")
}
