package redlot

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb/util"
)

func encodeKvKey(key []byte) (buf []byte) {
	buf = append(buf, typeKV)
	buf = append(buf, key...)
	return
}

func decodeKvKey(buf []byte) (key []byte) {
	if len(buf) < 4 {
		return nil
	}
	return buf[1:]
}

func Get(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return "", errNosArgs
	}

	v, err := db.Get(encodeKvKey(args[0]), nil)
	return string(v), err
}

func Set(args [][]byte) (interface{}, error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	fmt.Printf("SET %s %s\n", args[0], args[1])
	return nil, db.Put(encodeKvKey(args[0]), args[1], nil)
}

func Del(args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}

	return nil, db.Delete(encodeKvKey(args[0]), nil)
}

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

func Keys(args [][]byte) ([]string, error) {
	if len(args) < 3 {
		return []string{}, errNosArgs
	}

	ks := encodeKvKey(args[0])
	ke := encodeKvKey(args[1])
	limit, _ := strconv.Atoi(string(args[2]))

	var keys []string
	iter := db.NewIterator(&util.Range{Start: ks, Limit: ke}, nil)
	for iter.Next() {
		k := decodeKvKey(iter.Key())
		keys = append(keys, string(k))
		limit--
		if limit == 0 {
			break
		}
	}
	iter.Release()
	err := iter.Error()
	return keys, err
}

func Scan(args [][]byte) ([]string, error) {
	if len(args) < 3 {
		return []string{}, errNosArgs
	}

	ks := encodeKvKey(args[0])
	ke := encodeKvKey(args[1])
	limit, _ := strconv.Atoi(string(args[2]))

	var ret []string
	iter := db.NewIterator(&util.Range{Start: ks, Limit: ke}, nil)
	for iter.Next() {
		k := decodeKvKey(iter.Key())
		ret = append(ret, string(k))
		ret = append(ret, string(iter.Value()))
		limit--
		if limit == 0 {
			break
		}
	}
	iter.Release()
	err := iter.Error()
	return ret, err
}
