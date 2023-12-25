package gen

import (
	"go-server-gen/gen/conf"
	"go-server-gen/gen/data"
	"go-server-gen/gen/parse"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestGen(t *testing.T) {
	layout, idl, err := conf.GetConfig(LayoutYaml, IdlYaml)
	if err != nil {
		println("yaml err: ", err.Error())
		return
	}

	groups, _, err := data.ConfigToData(layout, idl)
	if err != nil {
		println("config to data err: ", err.Error())
		return
	}

	res, err := parse.GenServiceCode(layout, groups)
	if err != nil {
		println("gen code err: ", err.Error())
		return
	}
	for _, v := range res {
		println(v.File, "\n"+v.Code)
	}
}

func TestTemp(t *testing.T) {
	tmplStr := `{{if hasPrefix .Var "middleware"}}{{.Var}}{{else}}controller.{{.Var}}{{end}}`
	tmpl := template.Must(template.New("mytemplate").Funcs(template.FuncMap{
		"hasPrefix": strings.HasPrefix,
	}).Parse(tmplStr))

	data := struct {
		Var string
	}{
		Var: "AdminAuth",
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
