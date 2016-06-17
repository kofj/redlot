package redlot

func encodeHashKey(name, key []byte) (buf []byte) {
	buf = append(buf, typeHASH)
	buf = append(buf, uint32ToBytes(uint32(len(name)))...)
	buf = append(buf, name...)
	buf = append(buf, uint32ToBytes(uint32(len(key)))...)
	buf = append(buf, key...)
	return
}

func decodeHashKey(b []byte) (name, key []byte) {
	nameLen := bytesToUint32(b[1:5])
	name = b[5 : 5+nameLen]
	key = b[9+nameLen:]
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

// Hset will set a hashmap value by the key.
// Args: name string, key string, value any
func Hset(args [][]byte) (r interface{}, err error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}

	return
}

// Hset will return a hashmap value by the key.
// Args: name string, key string
func Hget(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return
}

// Hdel will delete a hashmap value by the key.
// Args: name string, key string
func Hdel(args [][]byte) (r interface{}, err error) {
	if len(args) < 2 {
		return nil, errNosArgs
	}

	return
}

// Hdel will incr a hashmap value by the key.
// Args: name string, key string, value int
func Hincr(args [][]byte) (r interface{}, err error) {
	if len(args) < 3 {
		return nil, errNosArgs
	}

	return
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

	return
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

// Hkeys will list the hashmap keys in the range.
// Args: name string, start string, end string, limit int
func Hkeys(args [][]byte) (r []string, err error) {
	if len(args) < 4 {
		return nil, errNosArgs
	}

	return
}

// Hgetall will list all keys/value in the hashmap.
// Args: name string
func Hgetall(args [][]byte) (r []string, err error) {
	if len(args) < 1 {
		return nil, errNosArgs
	}

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