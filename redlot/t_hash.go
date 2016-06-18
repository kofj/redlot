package redlot

import (
	"strconv"

	"github.com/syndtr/goleveldb/leveldb/util"
)

func encodeHashKey(name, key []byte) (buf []byte) {
	buf = append(buf, typeHASH)
	buf = append(buf, uint32ToBytes(uint32(len(name)))...)
	buf = append(buf, name...)
	buf = append(buf, key...)
	return
}

func decodeHashKey(b []byte) (name, key []byte) {
	nameLen := bytesToUint32(b[1:5])
	name = b[5 : 5+nameLen]
	key = b[5+nameLen:]
	return
}

func encodeHsizeKey(name []byte) (buf []byte) {
	buf = append(buf, typeHSIZE)
	buf = append(buf, uint32ToBytes(uint32(len(name)))...)
	buf = append(buf, name...)
	return
}

func decodeHsizeKey(b []byte) (key []byte) {
	return b[5:]
}

func hashSizeIncr(name []byte, incr int) {
	hsize := encodeHsizeKey(name)

	var size uint32
	if b, err := db.Get(hsize, nil); err == nil {
		size = bytesToUint32(b)
	}

	if incr > 0 {
		size += uint32(incr)
	} else {
		size = size - uint32(0-incr)
	}
	db.Put(hsize, uint32ToBytes(size), nil)
}

// Hset will set a hashmap value by the key.
// Args: name string, key string, value any
func Hset(args [][]byte) (r interface{}, err error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}

	hashSizeIncr(args[0], 1)
	err = db.Put(encodeHashKey(args[0], args[1]), args[2], nil)
	if err != nil {
		hashSizeIncr(args[0], -1)
		return nil, err
	}

	return
}

// Hset will return a hashmap value by the key.
// Args: name string, key string
func Hget(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}
	var b []byte
	b, err = db.Get(encodeHashKey(args[0], args[1]), nil)
	return string(b), err
}

// Hdel will delete a hashmap value by the key.
// Args: name string, key string
func Hdel(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return
}

func hincr(name, key []byte, increment int) (r int64, err error) {
	hash := encodeHashKey(name, key)
	v, _ := db.Get(hash, nil)
	var number int
	var exists bool
	if len(v) != 0 {
		number, err = strconv.Atoi(string(v))
		exists = true
		if err != nil {
			return -1, errNotInt
		}
	}
	number += increment
	r = int64(number)
	err = db.Put(hash, []byte(strconv.Itoa(number)), nil)
	if err == nil && !exists {
		hashSizeIncr(name, 1)
	}
	return
}

// Hincr will incr a hashmap value by the key.
// Args: name string, key string
func Hincr(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return hincr(args[0], args[1], 1)
}

// Hincrby will incr number a hashmap value by the key.
// Args: name string, key string, value int
func Hincrby(args [][]byte) (r interface{}, err error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}
	i, e := strconv.Atoi(string(args[2]))
	if e != nil {
		return -1, errNotInt
	}

	return hincr(args[0], args[1], i)
}

// Hexists will check the hashmap key is exists.
// Args: name string, key string
func Hexists(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return
}

// Hsize will return the hashmap size.
// Args: name string
func Hsize(args [][]byte) (r interface{}, err error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}

	var size int64
	if b, e := db.Get(encodeHsizeKey(args[0]), nil); e != nil {
		size = -1
	} else {
		size = int64(bytesToUint32(b))
	}

	return size, nil
}

// Hlist will list all hashmap in the range.
// Args: start string, end string, limit int
func Hlist(args [][]byte) (r []string, err error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}

	return
}

// Hrlist will reverse list all hashmap in the range.
// Args: start string, end string, limit int
func Hrlist(args [][]byte) (r []string, err error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}

	return
}

func hekys(args [][]byte, reverse bool) (r []string, err error) {
	if _, err = db.Get(encodeHsizeKey(args[0]), nil); err != nil {
		return
	}

	if len(args[1]) != 0 && string(args[1]) >= string(args[2]) {
		return []string{""}, nil
	}

	var ks, ke []byte
	if len(args[1]) == 0 {
		ks = append(ks, typeHASH)
		ks = append(ks, uint32ToBytes(uint32(len(args[0])))...)
		ks = append(ks, args[0]...)
	} else {
		ks = encodeHashKey(args[0], args[1])
	}

	if len(args[2]) == 0 {
		ke = append(ks, []byte{0xff}...)
	} else {
		ke = encodeHashKey(args[0], args[2])
	}

	limit, _ := strconv.Atoi(string(args[3]))

	iter := db.NewIterator(&util.Range{Start: ks, Limit: ke}, nil)
	if reverse {
		iter.Seek(ke)
		for iter.Prev() {
			_, key := decodeHashKey(iter.Key())
			r = append(r, string(key))
			limit--
			if limit <= 0 {
				break
			}
		}
	} else {
		for iter.Next() {
			_, key := decodeHashKey(iter.Key())
			r = append(r, string(key))
			limit--
			if limit <= 0 {
				break
			}
		}
	}
	iter.Release()
	err = iter.Error()

	return
}

// Hkeys will list the hashmap keys in the range.
// Args: name string, start string, end string, limit int
func Hkeys(args [][]byte) (r []string, err error) {
	if len(args) < 4 {
		return nil, errNosArgs
	}

	r, err = hekys(args, false)

	return
}

// Hrkeys will reverse list the hashmap keys in the range.
// Args: name string, start string, end string, limit int
func Hrkeys(args [][]byte) (r []string, err error) {
	if len(args) < 4 {
		return nil, errNosArgs
	}

	r, err = hekys(args, true)

	return
}

// Hgetall will list all keys/value in the hashmap.
// Args: name string
func Hgetall(args [][]byte) (r []string, err error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}

	if _, err = db.Get(encodeHsizeKey(args[0]), nil); err != nil {
		return
	}

	var buf []byte
	buf = append(buf, typeHASH)
	buf = append(buf, uint32ToBytes(uint32(len(args[0])))...)
	buf = append(buf, args[0]...)
	ke := append(buf, []byte{0xff}...)

	iter := db.NewIterator(&util.Range{Start: buf, Limit: ke}, nil)
	for iter.Next() {
		_, key := decodeHashKey(iter.Key())
		r = append(r, string(key))
		r = append(r, string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	return
}

// Hscan will list keys/value of the hashmap in the range.
// Args: name string, start string, end string, limit int
func Hscan(args [][]byte) (r []string, err error) {
	if len(args) < 4 {
		return nil, errNosArgs
	}

	return
}

// Hrscan will reverse list keys/value of the hashmap in the range.
// Args: name string, start string, end string, limit int
func Hrscan(args [][]byte) (r []string, err error) {
	if len(args) < 4 {
		return nil, errNosArgs
	}

	return
}

// Hclear will remove all value in the hashmap.
// Args: name string
func Hclear(args [][]byte) (r interface{}, err error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}

	return
}

// MultiHget will return multi hashmap value by keys.
func MultiHget(args [][]byte) (r []string, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return
}

// MultiHset will set multi hashmap value by keys.
func MultiHset(args [][]byte) (r interface{}, err error) {
	if len(args) < 3 && len(args)%2 == 0 {
		return nil, errNosArgs
	}

	return
}

// MultiHdel will delete multi hashmap value by keys.
func MultiHdel(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return
}
