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
