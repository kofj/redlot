package client

import "net"

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
	return
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Client) send(args []interface{}) (err error) {
	return
}

func (c *Client) recv() (r *Reply) {
	return
}
