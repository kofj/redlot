package client

import (
	"bytes"
	"os"
	"testing"
	"time"

	"../net"
	"../redlot"
)

var (
	client *Client
	err    error
)

func TestMain(m *testing.M) {
	// clean env
	os.RemoveAll("/tmp/data")
	os.RemoveAll("/tmp/meta")

	os.Exit(func() (r int) {
		options := &redlot.Options{
			DataPath: "/tmp",
		}

		go net.Serve(":9999", options)

		// Wait 1ms to start server.
		time.Sleep(5e6)

		r = m.Run()

		client.Close()
		os.RemoveAll("/tmp/data")
		os.RemoveAll("/tmp/meta")
		return r
	}())
}

func TestNewClient(t *testing.T) {
	o := &Options{
		Addr: "127.0.0.1:9999",
	}
	client, err = NewClient(o)
	if err != nil || client == nil {
		t.Logf("client: %+v, err: %v\n", client, err)
		t.Fail()
	}
}

func TestCmd(t *testing.T) {
	r := client.Cmd("set", "k", "v")
	if r.State != ReplyOK {
		t.Logf("Cmd [set k v] reply error: %s", r.State)
		t.Fail()
	}
	r = client.Cmd("get", "k")
	if r.State != ReplyOK {
		t.Logf("Cmd [get k] reply state error: %s", r.State)
		t.Fail()
	}
	if len(r.Data) != 1 {
		t.Logf("Cmd [get k] reply length error, expect 1, but %d", len(r.Data))
		t.Fail()
	}
	if string(r.Data[0]) != "v" {
		t.Logf("Cmd [get k] reply data error, expect string \"v\" , but %s", string(r.Data[0]))
		t.Fail()
	}

}

func TestSendBuf(t *testing.T) {
	var tests = []struct {
		in  []interface{}
		out string
	}{
		{[]interface{}{"set", "age", "19"}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},

		{[]interface{}{"set", "age", int(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", int8(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", int16(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", int32(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", int64(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},

		{[]interface{}{"set", "out", int(-1)}, "*3\r\n$3\r\nset\r\n$3\r\nout\r\n$2\r\n-1\r\n"},
		{[]interface{}{"set", "out", int8(-1)}, "*3\r\n$3\r\nset\r\n$3\r\nout\r\n$2\r\n-1\r\n"},
		{[]interface{}{"set", "out", int16(-1)}, "*3\r\n$3\r\nset\r\n$3\r\nout\r\n$2\r\n-1\r\n"},
		{[]interface{}{"set", "out", int32(-1)}, "*3\r\n$3\r\nset\r\n$3\r\nout\r\n$2\r\n-1\r\n"},
		{[]interface{}{"set", "out", int64(-1)}, "*3\r\n$3\r\nset\r\n$3\r\nout\r\n$2\r\n-1\r\n"},

		{[]interface{}{"set", "age", uint(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", uint8(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", uint16(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", uint32(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},
		{[]interface{}{"set", "age", uint64(19)}, "*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n"},

		{[]interface{}{"set", "price", float32(9.9)}, "*3\r\n$3\r\nset\r\n$5\r\nprice\r\n$3\r\n9.9\r\n"},
		{[]interface{}{"set", "price", float64(9.9)}, "*3\r\n$3\r\nset\r\n$5\r\nprice\r\n$3\r\n9.9\r\n"},
		{[]interface{}{"set", "price", 9.90}, "*3\r\n$3\r\nset\r\n$5\r\nprice\r\n$3\r\n9.9\r\n"},
	}

	var buf []byte
	var err error

	for k, test := range tests {
		buf, err = client.sendBuf(test.in)
		if err != nil {
			t.Logf("%d [% #v] => [% #v], error: %s\n", k, test.in, buf, err.Error())
			t.Fail()
		}
		if !bytes.Equal(buf, []byte(test.out)) {
			t.Logf("%d [% #v] => [% #v], expect: [% #v]\n", k, test.in, string(buf), test.out)
			t.Fail()
		}
	}
}
