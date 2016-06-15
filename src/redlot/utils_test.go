package redlot

import (
	"bytes"
	"testing"
)

func TestUint32ToBytes(t *testing.T) {
	expect := []byte{0x00, 0x00, 0x03, 0xe8}
	b := uint32ToBytes(1000)
	if !bytes.Equal(expect, b) {
		t.Logf("expect: % #X,but get: % #X", expect, b)
		t.Fail()
	}
}
