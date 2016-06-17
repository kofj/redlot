package redlot

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
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
