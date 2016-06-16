package net

import (
	"io"
	"log"
	"net"
	"sync"

	"../redlot"
)

var counter struct {
	sync.RWMutex
	ConnCounter uint64
	TotalCalls  uint64
}

func init() {
	// Register commands.
	// system info
	REG("INFO", StatusReply, info)

	// KV type
	REG("GET", BulkReply, redlot.Get)
	REG("SET", StatusReply, redlot.Set)
	REG("DEL", StatusReply, redlot.Del)
	REG("EXISTS", IntReply, redlot.Exists)
	REG("SETX", StatusReply, redlot.Setx)
	REG("SETEX", StatusReply, redlot.Setx) // Alias of SETX
	REG("TTL", IntReply, redlot.TTL)
	REG("EXPIRE", IntReply, redlot.Expire)
	REGL("KEYS", ListReply, redlot.Keys)
	REGL("SCAN", ListReply, redlot.Scan)
	REGL("RSCAN", ListReply, redlot.Rscan)

}

func Serve(addr string, options *redlot.Options) {
	// Open LevelDB with options.
	redlot.Open(options)

	// Create sockets listener.
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
		counter.Lock()
		counter.ConnCounter++
		counter.Unlock()

		go func(c net.Conn) {
			for {
				req, err := newRequset(c)
				if err == io.EOF {
					break
				}
				if err != nil {
					continue
				}

				counter.Lock()
				counter.TotalCalls++
				counter.Unlock()

				r := run(req.Cmd, req.Args)
				r.WriteTo(c)
			}

			c.Close()
			counter.Lock()
			counter.ConnCounter--
			counter.Unlock()

		}(conn)

	}
}
