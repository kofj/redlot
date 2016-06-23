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

func TestString(t *testing.T) {
	r := client.Cmd("get", "k").String()

	if r != "v" {
		t.Logf("expect string [v], but get[ %s]\n", r)
		t.Fail()
	}

}
