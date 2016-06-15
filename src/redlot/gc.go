package redlot

import (
	"github.com/syndtr/goleveldb/leveldb/util"

	"time"
)

// garbage collector
func gc() {
	for {
		// fmt.Println("\n.................................")
		now := time.Now().UTC().Unix()

		// scan KV and delete expired paired.
		iter := meta.NewIterator(&util.Range{Start: []byte{0x6b, 0x00}, Limit: []byte{0x6b, 0xff}}, nil)
		for iter.Next() {
			key := iter.Key()
			life := int64(bytesToUint64(iter.Value()))
			if now > life {
				db.Delete(key, nil)
				meta.Delete(key, nil)
			}
			// fmt.Printf("key: %s (% #x),life: %d, live: %v\n", key, key, life, life >= now)
		}

		time.Sleep(5e9)
	}
}
