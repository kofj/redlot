package redlot

import (
	"errors"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func encodeKvKey(key []byte) (buf []byte) {
	buf = append(buf, typeKV)
	buf = append(buf, key...)
	return
}

func decodeKvKey(buf []byte) (key []byte) {
	if len(buf) < 2 {
		return []byte{}
	}
	return buf[1:]
}

// Get a value by a key.
// Args: key string
func Get(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return "", errNosArgs
	}

	v, err := db.Get(encodeKvKey(args[0]), nil)
	return string(v), err
}

// Set a value by a key.
// Args: key string, value any
func Set(args [][]byte) (interface{}, error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	// fmt.Printf("SET %s %s\n", args[0], args[1])
	return nil, db.Put(encodeKvKey(args[0]), args[1], nil)
}

// Incr 1.
// Args: key string
func Incr(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}

	key := encodeKvKey(args[0])
	v, _ := db.Get(key, nil)
	var number int
	if len(v) != 0 {
		var err error
		number, err = strconv.Atoi(string(v))
		if err != nil {
			return nil, errNotInt
		}
	}
	number++
	return int64(number), db.Put(key, []byte(strconv.Itoa(number)), nil)
}

// Del will delete a value by a key.
// Args: key string
func Del(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}
	return nil, db.Delete(encodeKvKey(args[0]), nil)
}

// Exists will check key is exists.
// Args: key string
func Exists(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return int64(-1), errNosArgs
	}

	ret, err := db.Has(encodeKvKey(args[0]), nil)
	if ret {
		return int64(1), err
	}
	return int64(0), err
}

// Expire the key after timeout.
// Args: key string, seconds int
func Expire(args [][]byte) (interface{}, error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	key := encodeKvKey(args[0])
	has, _ := db.Has(key, nil)
	if has {
		ttl := strToInt64(string(args[1]))
		if ttl < 1 {
			return nil, errors.New("TTL must > 0, you set to " + string(args[1]))
		}
		bs := uint64ToBytes(uint64(time.Now().UTC().Unix() + ttl))
		if meta.Put(key, bs, nil) == nil {
			return int64(1), nil
		}
	}
	return int64(0), nil
}

// Setx will set a value by the key and expire it after timeout.
// Args: key string, value any, seconds int
func Setx(args [][]byte) (interface{}, error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}
	key := encodeKvKey(args[0])

	ttl := strToInt64(string(args[2]))
	if ttl < 1 {
		return nil, errors.New("TTL must > 0, you set to " + string(args[2]))
	}
	bs := uint64ToBytes(uint64(time.Now().UTC().Unix() + ttl))
	meta.Put(key, bs, nil)

	return nil, db.Put(key, args[1], nil)
}

// TTL will return the lifetime of the key.
// Args: key string
func TTL(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return int64(-1), errNosArgs
	}

	key := encodeKvKey(args[0])
	b, _ := meta.Get(key, nil)
	if len(b) < 1 {
		return int64(-1), nil
	}
	ttl := int64(bytesToUint64(b)) - time.Now().UTC().Unix()
	if ttl < 0 {
		ttl = -1
		meta.Delete(key, nil)
		db.Delete(key, nil)
	}
	return ttl, nil
}

func keys(args [][]byte, reverse bool) ([]string, error) {
	if len(args) < 3 {
		return []string{}, errNosArgs
	}

	ks := encodeKvKey(args[0])
	ke := encodeKvKey(args[1])
	limit, _ := strconv.Atoi(string(args[2]))

	var keys []string
	iter := db.NewIterator(&util.Range{Start: ks, Limit: ke}, nil)
	if reverse {
		iter.Seek(ke)
		for iter.Prev() {
			k := decodeKvKey(iter.Key())
			keys = append(keys, string(k))
			limit--
			if limit == 0 {
				break
			}
		}
	} else {
		for iter.Next() {
			k := decodeKvKey(iter.Key())
			keys = append(keys, string(k))
			limit--
			if limit == 0 {
				break
			}
		}
	}
	iter.Release()
	err := iter.Error()
	return keys, err
}

// Keys will list keys in the range.
// Args: start_key string, end_key string, limit_number int
func Keys(args [][]byte) ([]string, error) {
	return keys(args, false)
}

// Rkeys will reverse list keys in the range.
// Args: start_key string, end_key string, limit_number int
func Rkeys(args [][]byte) ([]string, error) {
	return keys(args, true)
}

func scan(args [][]byte, reverse bool) ([]string, error) {
	if len(args) < 3 {
		return []string{}, errNosArgs
	}

	ks := encodeKvKey(args[0])
	ke := encodeKvKey(args[1])
	limit, _ := strconv.Atoi(string(args[2]))

	var ret []string
	iter := db.NewIterator(&util.Range{Start: ks, Limit: ke}, nil)
	if reverse {
		iter.Seek(ke)
		for iter.Prev() {
			k := decodeKvKey(iter.Key())
			ret = append(ret, string(k))
			ret = append(ret, string(iter.Value()))
			limit--
			if limit == 0 {
				break
			}
		}
	} else {
		for iter.Next() {
			k := decodeKvKey(iter.Key())
			ret = append(ret, string(k))
			ret = append(ret, string(iter.Value()))
			limit--
			if limit == 0 {
				break
			}
		}
	}
	iter.Release()
	err := iter.Error()
	return ret, err
}

// Scan will list KV pair that keys in the range.
// Args: start_key string, end_key string, limit_number int
func Scan(args [][]byte) ([]string, error) {
	return scan(args, false)
}

// Rscan will reverse list KV pair that keys in the range.
// Args: start_key string, end_key string, limit_number int
func Rscan(args [][]byte) ([]string, error) {
	return scan(args, true)
}

// MultiGet wil batch read data from db.
func MultiGet(args [][]byte) (r []string, err error) {
	if len(args) < 1 {
		return []string{}, errNosArgs
	}
	for _, key := range args {
		v, _ := db.Get(encodeKvKey(key), nil)
		r = append(r, string(key))
		r = append(r, string(v))
	}
	return
}

// MultiSet wil batch write data to db.
func MultiSet(args [][]byte) (interface{}, error) {
	if len(args) < 2 || len(args)%2 == 1 {
		return []string{}, errNosArgs
	}

	batch := new(leveldb.Batch)
	for i := 0; i < len(args); i += 2 {
		batch.Put(encodeKvKey(args[i]), args[i+1])
	}
	return nil, db.Write(batch, nil)
}
