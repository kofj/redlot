package redlot

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// clean env
	os.RemoveAll("/tmp/data")
	os.RemoveAll("/tmp/meta")

	os.Exit(func() (r int) {
		r = m.Run()
		os.RemoveAll("/tmp/data")
		os.RemoveAll("/tmp/meta")
		return r
	}())
}

func TestOpen(t *testing.T) {
	o := &Options{
		DataPath: "/tmp",
	}
	Open(o)
	if db == nil || meta == nil {
		t.Fail()
	}
}
