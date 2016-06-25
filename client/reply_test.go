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

func TestInt(t *testing.T) {
	r := client.Cmd("get", "k").Int()

	if r != 0 {
		t.Logf("expect int [0], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int", "9223372036854775807")
	r = client.Cmd("get", "int").Int()

	if r != 9223372036854775807 {
		t.Logf("expect int [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}
}
