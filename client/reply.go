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
	i, _ := strconv.Atoi(r.String())
	return i
}

func (r *Reply) Int8() int8 {
	i, _ := strconv.ParseInt(r.String(), 10, 8)
	return int8(i)
}

func (r *Reply) Int16() int16 {
	i, _ := strconv.ParseInt(r.String(), 10, 16)
	return int16(i)
}

func (r *Reply) Int32() int32 {
	i, _ := strconv.ParseInt(r.String(), 10, 32)
	return int32(i)
}

func (r *Reply) Int64() int64 {
	i, _ := strconv.ParseInt(r.String(), 10, 64)
	return i
}

func (r *Reply) Uint8() uint8 {
	i, _ := strconv.ParseUint(r.String(), 10, 8)
	return uint8(i)
}

func (r *Reply) Uint16() uint16 {
	i, _ := strconv.ParseUint(r.String(), 10, 16)
	return uint16(i)
}

func (r *Reply) Uint32() uint32 {
	i, _ := strconv.ParseUint(r.String(), 10, 32)
	return uint32(i)
}

func (r *Reply) Uint64() uint64 {
	i, _ := strconv.ParseUint(r.String(), 10, 64)
	return i
}

func (r *Reply) Uint() uint {
	return uint(r.Uint64())
}
