package redlot

import "fmt"

func encode_kv_key(key []byte) (buf []byte) {
	buf = append(buf, TYPE_KV)
	buf = append(buf, key...)
	buf = append(buf, Uint32ToBytes(uint32(len(key)))...)
	return
}

func decode_kv_key(buf []byte) (key []byte) {
	if len(buf) < 4 {
		return nil
	}
	return buf[1 : len(buf)-3]
}

func Get(args [][]byte) (interface{}, error) {
	v, err := db.Get(encode_kv_key(args[0]), nil)
	return string(v), err
}

func Set(args [][]byte) (interface{}, error) {
	fmt.Printf("SET %s %s\n", args[0], args[1])
	return nil, db.Put(encode_kv_key(args[0]), args[1], nil)
}

func Del(args [][]byte) (interface{}, error) {
	return nil, db.Delete(encode_kv_key(args[0]), nil)
}

func Exists(args [][]byte) (interface{}, error) {
	ret, err := db.Has(encode_kv_key(args[0]), nil)
	if ret {
		return int64(1), err
	}
	return int64(0), err
}
