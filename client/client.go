package client

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
)

type Client struct {
	conn net.Conn
}

func NewClient(o *Options) (*Client, error) {
	conn, err := o.getDialer()()
	return &Client{
		conn: conn,
	}, err
}

func (c *Client) Cmd(args ...interface{}) (r *Reply) {
	r = &Reply{
		State: ReplyError,
	}

	if c.conn == nil {
		r.State = ReplyFail
		return
	}

	if err := c.send(args); err != nil {
		r.State = ReplyFail
		return
	}
	r = c.recv()

	return
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Client) send(args []interface{}) (err error) {
	var buf []byte

	buf, err = c.send_buf(args)
	if err == nil {
		_, err = c.conn.Write(buf)
	}
	return
}

func (c *Client) recv() (r *Reply) {
	var line []byte
	reader := bufio.NewReader(c.conn)

	// *<number of arguments>CRLF
	line, err = reader.ReadBytes('\n')
	if err != nil {
		return
	}

	if line[0] == '+' || line[0] == '-' {
		return &Reply{
			State: string(line[1 : len(line)-2]),
		}
	}
	if line[0] == ':' {
		return &Reply{
			State: ReplyOK,
			Data:  [][]byte{line[1 : len(line)-2]},
		}
	}

	if line[0] == '$' {
		var dLen int
		dLen, err = strconv.Atoi(string(line[1 : len(line)-2]))
		if err != nil {
			return
		}
		line, _ = reader.ReadBytes('\n')

		return &Reply{
			State: ReplyOK,
			Data:  [][]byte{line[:dLen]},
		}
	}

	var data [][]byte
	if line[0] == '*' {
		var count int
		count, err = strconv.Atoi(string(line[1 : len(line)-2]))
		if err != nil {
			return
		}
		for i := 0; i < count; i++ {
			line, err = reader.ReadBytes('\n')
			if err != nil {
				return
			}
			dLen, _ := strconv.Atoi(string(line[1 : len(line)-2]))
			line, err = reader.ReadBytes('\n')
			if err != nil {
				return
			}
			data = append(data, line[:dLen])
		}
	}

	r = &Reply{
		State: ReplyOK,
		Data:  data,
	}

	return
}

func (c *Client) send_buf(args []interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("*%d\r\n", len(args)))

	for _, arg := range args {
		switch arg.(type) {
		case string:
			buf.WriteString(fmt.Sprintf("$%d\r\n", len(arg.(string))))
			buf.WriteString(arg.(string) + "\r\n")
			continue
		}
	}

	b = buf.Bytes()

	return
}
