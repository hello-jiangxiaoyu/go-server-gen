package gen

import (
	"go-server-gen/gen/code"
	"go-server-gen/gen/conf"
	"go-server-gen/gen/phase"
	"go-server-gen/utils"
	"regexp"
	"testing"
)

func TestGen(t *testing.T) {
	layout, idl, err := conf.GetConfig(LayoutYaml, IdlYaml)
	if err != nil {
		println("yaml err: ", err.Error())
		return
	}

	groups, _, err := phase.ConfigToCode(layout, idl)
	if err != nil {
		println("phase config err: ", err.Error())
		return
	}

	res, err := code.GenGroupCode(layout, groups)
	if err != nil {
		println("gen code err: ", err.Error())
		return
	}
	for _, v := range res {
		println(v.File, "\n"+v.Code)
	}
}

func TestTemplate(t *testing.T) {
	p := struct {
		Handlers map[string]string
	}{Handlers: map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}}
	const tpl = `
{{- range $K, $Handler := .}}
{{$K}} --> {{$Handler}}
{{- end }}

{{.a}}
`

	format, err := utils.PhaseTemplate(tpl, p.Handlers)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(format))
}

func TestTemp(t *testing.T) {
	src := []string{
		`func Handler(int a){}`,
		`func handler () {}`,
		`func (r *User)handler(){}`,
		`func (*User) handler(){}`,
		`func (*User) handler(a string, b int){}`,
		`func (*User) handler(a string, b int) {}`,
	}
	for _, h := range src {
		regFunction := regexp.MustCompile(`func .*handler\(.*\).*{`) // 获取go结构体uri tag
		fCode, err := utils.FormatCode([]byte(h))
		if err != nil {
			println("err: ", err.Error())
			continue
		}
		if regFunction.MatchString(string(fCode)) {
			println("Yes", string(fCode))
		} else {
			println("No")
		}
	}
}
