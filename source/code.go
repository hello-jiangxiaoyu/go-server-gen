package source

import (
	"go-server-gen/conf"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"strings"
)

type GlobalData struct {
	ProjectName string
	Pkg         map[string]conf.Package
}

// GenPackageCode 生成默认代码
func GenPackageCode(layout conf.LayoutConfig, resp ResponsePackage, code map[string]writer.WriteCode) error {
	projectName, err := utils.GetProjectName()
	if err != nil {
		return err
	}
	pkg := layout.Pkg
	pkg["Context"] = resp.Context
	pkg["Code"] = resp.Code
	pkg["Return"] = resp.Return
	pkg["ReturnType"] = resp.ReturnType
	pkg["BindCode"] = resp.BindCode
	pkg["ResponseCode"] = resp.ResponseCode
	pkg["HandleFunc"] = resp.HandleFunc
	for _, tpl := range layout.GlobalTemplate {
		file, body, err := utils.ParseSource(tpl.Path, tpl.Body, GlobalData{ProjectName: projectName, Pkg: pkg})
		if err != nil {
			return err
		}
		if tpl.FirstLine != "" && body != "" {
			lines := strings.SplitN(body, "\n", -1)
			lines[0] = tpl.FirstLine
			body = strings.Join(lines, "\n")
		}
		code[file] = writer.WriteCode{
			File:  file,
			Code:  body,
			Write: tpl.Write,
		}
	}
	return nil
}

// 从embed中读取文件内容
func GetEmbedContent(path string) string {
	res, err := code.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(res)
}
