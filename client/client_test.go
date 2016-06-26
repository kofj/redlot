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
	var args []interface{}
	var buf, expect []byte
	var err error

	// test string arg
	args = []interface{}{
		"set",
		"age",
		"19",
	}
	expect = []byte("*3\r\n$3\r\nset\r\n$3\r\nage\r\n$2\r\n19\r\n")

	buf, err = client.sendBuf(args)
	if err != nil {
		t.Logf("expect err is nil, but %s\n", err.Error())
		t.Fail()
	}
	if !bytes.Equal(buf, expect) {
		t.Logf("expect buf is [% #x], but get [% #x]", buf, expect)
		t.Fail()
	}
}
