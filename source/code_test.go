package source

import (
	"go-server-gen/conf"
	"go-server-gen/writer"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestCode(t *testing.T) {
	var layout conf.LayoutConfig
	if err := yaml.Unmarshal(serverTpl, &layout); err != nil {
		println(err.Error())
		t.FailNow()
	}

	res := make(map[string]writer.WriteCode)
	if err := GenPackageCode(layout, responseMap["gin"], res); err != nil {
		println(err.Error())
		t.FailNow()
	}

	if err := writer.Write(res); err != nil {
		println(err.Error())
		t.FailNow()
	}
}
