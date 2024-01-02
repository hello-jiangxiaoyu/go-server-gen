package source

import (
	"go-server-gen/conf"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"testing"
)

func TestCode(t *testing.T) {
	layout, err := conf.GetLayoutConfig("hertz", "")
	if err != nil {
		return
	}

	res := make(map[string]writer.WriteCode)
	if err = GenPackageCode(layout, "hertz", "zero", res); err != nil {
		println(err.Error())
		t.FailNow()
	}

	if err = writer.Write(res); err != nil {
		println(err.Error())
		t.FailNow()
	}
}

func TestDuplicateImport(t *testing.T) {
	println(utils.RemoveDuplicateImport(`
import (
"go-server-gen/conf"
	"go-server-gen/utils"
	"go-server-gen/writer"
"go-server-gen/writer"
	"testing"
	"gopkg.in/yaml.v3"
	"gopkg.in/yaml.v3"
)
asdf
as
df
adsa
sd
`))
}
