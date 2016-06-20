package client

const (
	ReplyOK       = "ok"
	ReplyNotFound = "not_found"
	ReplyError    = "error"
)

type Reply struct {
	State string
	Data  []byte
}
