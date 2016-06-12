package net

import "errors"

type CmdFunc func([][]byte) ([][]byte, error)
type REPLY_TYPE uint8

const (
	STATUS_REPLY REPLY_TYPE = 1 + iota
	ERR_REPLY
	INT_REPLY
	BULK_REPLY
	LIST_REPLY
)

var (
	cmdFuncs  = map[string]CmdFunc{}
	replyType = map[string]REPLY_TYPE{}
)

func REG(cmd string, types REPLY_TYPE, f CmdFunc) {
	cmdFuncs[cmd] = f
	replyType[cmd] = types
}

func RUN0(cmd string, args [][]byte) ([][]byte, error) {
	f, ok := cmdFuncs[cmd]
	if !ok {
		return nil, errors.New("unknwon command '" + cmd + "'")
	}
	return f(args)
}

func RUN(cmd string, args [][]byte) (reply Reply) {
	f, ok := cmdFuncs[cmd]
	if !ok {
		reply = &ErrReply{
			Msg: "unknwon command '" + cmd + "'",
		}
		return
	}

	t, ok2 := replyType[cmd]
	if !ok2 {
		reply = &ErrReply{
			Msg: "unknwon reply type of command '" + cmd + "'",
		}
		return
	}

	data, ferr := f(args)
	if ferr != nil {
		reply = &ErrReply{
			Msg: ferr.Error(),
		}
	}

	switch t {
	case STATUS_REPLY:
		reply = &StatusReply{
			Code: string(data[0]),
		}
		break
	case INT_REPLY:
		break

	case BULK_REPLY:
		break

	case LIST_REPLY:
		break
	}

	return
}
