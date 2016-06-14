package redlot

import (
	"log"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *leveldb.DB
)

// Open LevelDB.
func Open(o *Options) {
	o.DataPath = filepath.Clean(o.DataPath)
	if !filepath.IsAbs(o.DataPath) {
		log.Fatalf("[%s] not Abs path.", o.DataPath)
	}
	opts := o.convert()

	var err error
	db, err = leveldb.OpenFile(o.DataPath, opts)
	if err != nil {
		log.Fatalln("open db error:", err.Error())
	}
}

// Close LevelDB
func Close() {
	db.Close()
}
