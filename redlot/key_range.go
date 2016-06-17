package redlot

import "github.com/syndtr/goleveldb/leveldb/util"

// KeyRange will return key range.
func KeyRange() string {
	kr := "key_range.kv\n\t"
	iter := db.NewIterator(&util.Range{Start: []byte{0x6b, 0x00}}, nil)
	iter.Next()
	kr += "\"" + string(decodeKvKey(iter.Key())) + "\" - "

	iter.Last()
	iter.Prev()
	kr += "\"" + string(decodeKvKey(iter.Key())) + "\"\n"

	return kr
}
