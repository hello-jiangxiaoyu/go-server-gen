package gen

import (
	"go-server-gen/gen/conf"
	"go-server-gen/gen/phase"
	"os"
	"testing"
	"text/template"
)

func TestTemplate(tt *testing.T) {
	t := template.New("test")
	t = template.Must(t.Parse(conf.DefaultApiTemplate))
	p := phase.Api{
		Method:   "GET",
		Path:     "/v1/apps",
		Tag:      "app",
		Summary:  "获取app列表",
		FuncName: "GetAppList",
		ReqName:  "ReqApp",
		Context:  "*gin.Context",
		ReqParam: []phase.Param{{
			Name:        "dev",
			From:        "query",
			Type:        "string",
			Required:    "false",
			Description: "是否为开发阶段",
		}},
	}
	_ = t.Execute(os.Stdout, p)
}

func TestApi(tt *testing.T) {
	input := `type LoginReq  struct {
		UserName  string
		Password  string
	}
	type  UserReq struct{
		UserID int
	}`

	message, err := phase.SplitMessage(input)
	if err != nil {
		return
	}
	for k, v := range message {
		println(k, ":", v)
	}
}
