package source

import (
	"go-server-gen/conf"
	"go-server-gen/writer"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestEmbed(t *testing.T) {
	src, err := code.ReadFile("internal/gin/bind.go")
	if err != nil {
		println(err.Error())
		t.FailNow()
	}
	println(string(src))
}

func TestCode(t *testing.T) {
	var layout conf.LayoutConfig
	if err := yaml.Unmarshal(serverTpl, &layout); err != nil {
		println(err.Error())
		t.FailNow()
	}

	code := make(map[string]writer.WriteCode)
	if err := GenPackageCode(layout, responseMap["fiber"], code); err != nil {
		println(err.Error())
		t.FailNow()
	}

	if err := writer.Write(code); err != nil {
		println(err.Error())
		t.FailNow()
	}
}
