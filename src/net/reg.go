package net

type CmdFunc func([][]byte) (interface{}, error)
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
		if data == nil {
			reply = &StatusReply{
				Code: "OK",
			}
		} else {
			reply = &StatusReply{
				Code: (data.(string)),
			}
		}
		break
	case INT_REPLY:
		reply = &IntReply{
			Nos: data.(int64),
		}
		break

	case BULK_REPLY:
		break

	case LIST_REPLY:
		break
	}

	return
}
