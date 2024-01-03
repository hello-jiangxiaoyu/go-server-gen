package gen

import (
	"testing"
)

func TestGen(t *testing.T) {
	if err := ExecuteUpdate("fiber", "", "../conf/test-idl.yaml"); err != nil {
		println(err.Error())
		t.FailNow()
	}
	println("Success")
}

func TestTemp(t *testing.T) {

}
