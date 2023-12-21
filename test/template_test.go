package test

import (
	"os"
	"testing"
	"text/template"
)

type Para struct {
	Name        string
	From        string
	Type        string
	Required    string
	Description string
}
type Api struct {
	Method     string
	Path       string
	Summary    string
	FuncName   string
	ReqName    string
	Tag        string
	Context    string
	ReqPara    []Para
	Middleware []string
}
type ApiGroup struct {
	Apis       []Api
	Import     []string
	Middleware []string
}

const tpl = `
// {{.FuncName}}
// @Tags	{{.Tag}}
// @Summary	{{.Summary}}
{{- range $_, $Para := .ReqPara}}
// @Param	{{$Para.Name}}	{{$Para.From}}	{{$Para.Type}}	{{$Para.Required}}	"{{$Para.Description}}"
{{- end }}
// @Success	200
// @Router	{{.Path}} [{{.Method}}]
func {{.FuncName}}(c {{.Context}}) {
	// you controller code hear
}
`

func TestTemplate(tt *testing.T) {
	t := template.New("test")
	t = template.Must(t.Parse(tpl))
	p := Api{
		Method:   "GET",
		Path:     "/v1/apps",
		Tag:      "app",
		Summary:  "获取app列表",
		FuncName: "GetAppList",
		ReqName:  "ReqApp",
		Context:  "*gin.Context",
		ReqPara: []Para{{
			Name:        "dev",
			From:        "query",
			Type:        "string",
			Required:    "false",
			Description: "是否为开发阶段",
		}},
	}
	_ = t.Execute(os.Stdout, p)
}
