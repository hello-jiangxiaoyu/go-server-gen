package source

import (
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/parse"
	"go-server-gen/writer"
	"html/template"
	"os"
	"testing"
)

var (
	serverType = "gin"
	logType    = "zero"
)

func TestNewCode(t *testing.T) {
	layout, err := conf.GetLayoutConfig(serverType, logType, "")
	if err != nil {
		t.FailNow()
	}

	res, err := GenPackageCode(layout, "", false)
	if err != nil {
		t.FailNow()
	}

	for _, r := range res {
		println(r.File + "\n" + r.Code)
	}
}

func TestUpdateCode(t *testing.T) {
	layout, idl, err := conf.GetConfig(serverType, logType, "", "test-idl.yaml")
	if err != nil {
		t.FailNow()
	}

	services, _, err := data.ConfigToData(layout, idl)
	if err != nil {
		t.FailNow()
	}

	res := make(map[string]writer.WriteCode)
	if err = parse.GenServiceCode(layout, services, res); err != nil {
		t.FailNow()
	}

	if writer.Write(res) != nil {
		t.FailNow()
	}
}

func TestTemp(t *testing.T) {
	item := map[string]any{"Name": "zhangsan"}
	tmpl := template.Must(template.New("test").Parse(`Hello {{.Name}}, this is a string: {{printf "{{item.name}}"}}`))
	_ = tmpl.Execute(os.Stdout, item) // 原样输出字符串{{item.name}}
}
