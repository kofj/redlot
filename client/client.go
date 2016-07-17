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

// NewClient will connect to the server and create a client.
func NewClient(o *Options) (*Client, error) {
	conn, err := o.getDialer()()
	return &Client{
		conn: conn,
	}, err
}

// Cmd will send command, receive data from server and build reply.
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

// Close socks.
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Client) send(args []interface{}) (err error) {
	var buf []byte

	buf, err = c.sendBuf(args)
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

func (c *Client) sendBuf(args []interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	var s, size string
	buf.WriteString(fmt.Sprintf("*%d\r\n", len(args)))

	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			s = t

		case int:
			s = strconv.FormatInt(int64(t), 10)
		case int8:
			s = strconv.FormatInt(int64(t), 10)
		case int16:
			s = strconv.FormatInt(int64(t), 10)
		case int32:
			s = strconv.FormatInt(int64(t), 10)
		case int64:
			s = strconv.FormatInt(int64(t), 10)

		case uint:
			s = strconv.FormatUint(uint64(t), 10)
		case uint8:
			s = strconv.FormatUint(uint64(t), 10)
		case uint16:
			s = strconv.FormatUint(uint64(t), 10)
		case uint32:
			s = strconv.FormatUint(uint64(t), 10)
		case uint64:
			s = strconv.FormatUint(uint64(t), 10)

		case float32:
			s = strconv.FormatFloat(float64(t), 'f', -1, 32)

		case float64:
			s = strconv.FormatFloat(t, 'f', -1, 64)

		case bool:
		case nil:
		case []byte:
		case [][]byte:
		case []string:

		}

		size = fmt.Sprintf("$%d\r\n", len(s))
		buf.WriteString(size)
		buf.WriteString(s + "\r\n")

	}

	b = buf.Bytes()

	return
}
