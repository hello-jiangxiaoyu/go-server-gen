package gen

import (
	"go-server-gen/gen/conf"
	"go-server-gen/gen/data"
	"go-server-gen/gen/parse"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	layout, idl, err := conf.GetConfig(LayoutYaml, IdlYaml)
	if err != nil {
		println("yaml err: ", err.Error())
		os.Exit(1)
	}

	groups, _, err := data.ConfigToData(layout, idl)
	if err != nil {
		println("config to data err: ", err.Error())
		os.Exit(1)
	}

	res, err := parse.GenServiceCode(layout, groups)
	if err != nil {
		println("gen code err: ", err.Error())
		os.Exit(1)
	}
	for _, v := range res {
		println(v.File, "\n"+v.Code)
	}
}

func TestTemp(t *testing.T) {
	yamlStr := `
a: |
  This is a
    multi-line
      string.
b: |-
  This is a
    multi-line
      string.
c: >
  This is a
    multi-line
      string.
`
	res := make(map[string]string)
	if err := yaml.Unmarshal([]byte(yamlStr), &res); err != nil {
		return
	}
	for _, v := range res {
		println(v)
	}
}
