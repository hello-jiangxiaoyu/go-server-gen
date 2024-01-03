package writer

import "testing"

func TestWriter(t *testing.T) {
	src := []byte("ddd")
	if err := writeFile("hello/world.txt", src, false); err != nil {
		println(err.Error())
		t.FailNow()
	}
}
