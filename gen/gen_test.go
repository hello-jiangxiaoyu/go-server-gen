package gen

import (
	"go-server-gen/gen/code"
	"go-server-gen/gen/conf"
	"go-server-gen/gen/phase"
	"testing"
)

func TestGen(t *testing.T) {
	layout, idl, err := conf.GetConfig(LayoutYaml, IdlYaml)
	if err != nil {
		println("yaml: ", err.Error())
		return
	}

	groups, err := phase.ConfigToCode(layout, idl)
	if err != nil {
		return
	}

	for _, g := range groups {
		_, err = code.GenGroupCode(g)
		if err != nil {
			return
		}
	}
}
