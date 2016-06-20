package client

import (
	"net"
	"time"
)

type Options struct {
	Network     string
	Addr        string
	Dialer      func() (net.Conn, error)
	DialTimeout time.Duration
}

func (o *Options) getDialer() func() (net.Conn, error) {
	if o.Dialer != nil {
		return o.Dialer
	}
	return func() (net.Conn, error) {
		return net.DialTimeout(o.getNetwork(), o.Addr, o.getDialTimeout())
	}
}

func (o *Options) getNetwork() string {
	if o.Network == "" {
		return "tcp"
	}

	return o.Network
}

func (o *Options) getDialTimeout() time.Duration {
	if o.DialTimeout == 0 {
		return 5 * time.Second
	}
	return o.DialTimeout
}
