package source

import (
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/parse"
	"go-server-gen/writer"
	"testing"
)

var (
	serverType = "gin"
	logType    = "zero"
)

func TestNewCode(t *testing.T) {
	layout, err := conf.GetLayoutConfig(serverType, "")
	if err != nil {
		t.FailNow()
	}

	res := make(map[string]writer.WriteCode)
	if err = GenPackageCode(layout, serverType, logType, res); err != nil {
		t.FailNow()
	}

	for _, r := range res {
		println(r.File + "\n" + r.Code)
	}
}

func TestUpdateCode(t *testing.T) {
	layout, idl, err := conf.GetConfig(serverType, "", "test-idl.yaml")
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

	for _, r := range res {
		println(r.File + "\n" + r.Code)
	}
}
