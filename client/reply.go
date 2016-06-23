package client

import "strconv"

const (
	ReplyOK       = "OK"
	ReplyNotFound = "not_found"
	ReplyError    = "error"
	ReplyFail     = "fail"
)

type Reply struct {
	State string
	Data  [][]byte
}

func (r *Reply) bytes() []byte {
	if len(r.Data) > 0 {
		return r.Data[0]
	}

	return []byte{}
}

func (r *Reply) Bytes() []byte {
	return r.bytes()
}

func (r *Reply) String() string {
	return string(r.bytes())
}

func (r *Reply) Int() int {
	i, err := strconv.Atoi(r.String())
	if err != nil {
		return 0
	}
	return i
}
