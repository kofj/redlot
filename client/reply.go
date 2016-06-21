package client

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
