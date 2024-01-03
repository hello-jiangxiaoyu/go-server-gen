package source

import (
	"go-server-gen/conf"
	"go-server-gen/gen"
	"go-server-gen/writer"
	"testing"
)

var server = "gin"

func TestNewCode(t *testing.T) {
	layout, err := conf.GetLayoutConfig(server, "")
	if err != nil {
		return
	}

	res := make(map[string]writer.WriteCode)
	if err = GenPackageCode(layout, server, "zero", res); err != nil {
		println(err.Error())
		t.FailNow()
	}

	if err = writer.Write(res); err != nil {
		println(err.Error())
		t.FailNow()
	}
}

func TestUpdateCode(t *testing.T) {
	if err := gen.ExecuteUpdate(server, "", "../conf/test-idl.yaml"); err != nil {
		println(err.Error())
		t.FailNow()
	}
	println("Success")
}
