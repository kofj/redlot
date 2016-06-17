package net

type cmdFunc func([][]byte) (interface{}, error)
type listCmdFunc func([][]byte) ([]string, error)
type replyType uint8

const (
	// StatusReply etc defined reply type.
	StatusReply replyType = 1 + iota
	ErrReply
	IntReply
	BulkReply
	ListReply
)

var (
	cmdFuncs     = map[string]cmdFunc{}
	listCmdFuncs = map[string]listCmdFunc{}
	replyTypes   = map[string]replyType{}
)

// REG will register a cmd function to server.
func REG(cmd string, types replyType, f cmdFunc) {
	cmdFuncs[cmd] = f
	replyTypes[cmd] = types
}

// REGL will regiter a list cmd function to server.
func REGL(cmd string, types replyType, f listCmdFunc) {
	listCmdFuncs[cmd] = f
}

// Execute cmd function and generate reply.
func run(cmd string, args [][]byte) (r reply) {
	f, ok := cmdFuncs[cmd]
	fl, okl := listCmdFuncs[cmd]
	if !ok && !okl {
		r = &errReply{
			Msg: "unknwon command '" + cmd + "'",
		}
		return
	}
	if okl {
		data, ferr := fl(args)

		if ferr != nil {
			return &errReply{
				Msg: ferr.Error(),
			}
		}
		return &listReply{
			List: data,
		}
	}

	t, ok2 := replyTypes[cmd]
	if !ok2 {
		r = &errReply{
			Msg: "unknwon reply type of command '" + cmd + "'",
		}
		return
	}

	data, ferr := f(args)

	if ferr != nil {
		return &errReply{
			Msg: ferr.Error(),
		}
	}

	switch t {
	case StatusReply:
		if data == nil {
			r = &statusReply{
				Code: "OK",
			}
		} else {
			r = &statusReply{
				Code: (data.(string)),
			}
		}
		break
	case IntReply:
		r = &intReply{
			Nos: data.(int64),
		}
		break

	case BulkReply:
		r = &bulkReply{
			Bulk: data.(string),
		}
		break

		// case ListReply:
		// 	break
	}

	return r
}
