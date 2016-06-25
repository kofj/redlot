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

	// test value out of range
	client.Cmd("set", "int", "9223372036854775808")
	r = client.Cmd("get", "int").Int()

	if r != 9223372036854775807 {
		t.Logf("expect int [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestInt32(t *testing.T) {
	client.Cmd("set", "int32", "2147483647")
	r := client.Cmd("get", "int32").Int32()

	if r != 2147483647 {
		t.Logf("expect int [2147483647], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int32", "-2147483648")
	r = client.Cmd("get", "int32").Int32()

	if r != -2147483648 {
		t.Logf("expect int [-2147483648], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int32", "2147483648")
	r = client.Cmd("get", "int32").Int32()

	if r != 2147483647 {
		t.Logf("expect int [2147483647], but get[ %d]\n", r)
		t.Fail()
	}
}

func TestInt64(t *testing.T) {
	client.Cmd("set", "int64", "9223372036854775807")
	r := client.Cmd("get", "int64").Int64()

	if r != 9223372036854775807 {
		t.Logf("expect int [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int64", "-9223372036854775808")
	r = client.Cmd("get", "int64").Int64()

	if r != -9223372036854775808 {
		t.Logf("expect int [-9223372036854775808], but get[ %d]\n", r)
		t.Fail()
	}

	client.Cmd("set", "int64", "9223372036854775808")
	r = client.Cmd("get", "int64").Int64()

	if r != 9223372036854775807 {
		t.Logf("expect int [9223372036854775807], but get[ %d]\n", r)
		t.Fail()
	}
}
