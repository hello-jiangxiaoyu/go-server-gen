package gen

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGen(t *testing.T) {
	if err := ExecuteUpdate("fiber"); err != nil {
		println(err.Error())
		t.FailNow()
	}
	println("Success")
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
