package client

import (
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
		os.RemoveAll("/tmp/data")
		os.RemoveAll("/tmp/meta")
		return r
	}())
}
