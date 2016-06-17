package redlot

import (
	"encoding/binary"
	"strconv"
)

func uint32ToBytes(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

func uint64ToBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func bytesToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func strToInt64(str string) int64 {
	u, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		u = 0
	}
	return u
}
