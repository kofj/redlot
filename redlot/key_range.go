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

	kr += "hash_range.kv\n\t"
	iter = db.NewIterator(&util.Range{Start: []byte{0x48, 0x00}}, nil)
	iter.Next()
	kr += "\"" + string(decodeHsizeKey(iter.Key())) + "\" - "

	iter.Seek([]byte{0x48, 0xff})
	iter.Prev()
	kr += "\"" + string(decodeHsizeKey(iter.Key())) + "\"\n"

	return kr
}
