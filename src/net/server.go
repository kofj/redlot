package net

import (
	"io"
	"log"
	"net"
	"sync"

	"../redlot"
)

var info struct {
	sync.RWMutex
	ConnCounter uint64
	TotalCalls  uint64
	reply       Reply
}

func init() {
	REG("INFO", STATUS_REPLY, Info)
	REG("GET", BULK_REPLY, redlot.Get)
	REG("SET", STATUS_REPLY, redlot.Set)
	REG("DEL", STATUS_REPLY, redlot.Del)
	REG("EXISTS", INT_REPLY, redlot.Exists)
}

func Serve(addr string, options *redlot.Options) {
	redlot.Open(options)

	l, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalf("Listen error: %v\n", err.Error())
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Wait for a connection error: %s\n", err.Error())
		}

		// Count connecion
		info.Lock()
		info.ConnCounter++
		info.Unlock()

		go func(c net.Conn) {
			for {
				req, err := newRequset(c)
				if err == io.EOF {
					break
				}
				if err != nil {
					continue
				}

				info.Lock()
				info.TotalCalls++
				info.Unlock()

				reply := RUN(req.Cmd, req.Args)
				reply.WriteTo(c)
			}

			c.Close()
			info.Lock()
			info.ConnCounter--
			info.Unlock()

		}(conn)

	}
}
