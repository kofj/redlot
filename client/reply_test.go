package client

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	r := client.Cmd("get", "k").Bytes()

	if !bytes.Equal(r, []byte("v")) {
		t.Logf("expect bytes [% #x], but get[ % #x]\n", []byte("v"), r)
		t.Fail()
	}
}
