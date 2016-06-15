package redlot

import (
	"log"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db       *leveldb.DB
	meta     *leveldb.DB
	metaPath string
	dataPath string
)

// Open LevelDB.
func Open(o *Options) {
	o.DataPath = filepath.Clean(o.DataPath)
	if !filepath.IsAbs(o.DataPath) {
		log.Fatalf("[%s] not Abs path.", o.DataPath)
	}
	metaPath = filepath.Join(o.DataPath, "meta")
	dataPath = filepath.Join(o.DataPath, "data")
	opts := o.convert()

	var err error
	meta, err = leveldb.OpenFile(metaPath, opts)
	if err != nil {
		log.Fatalln("open meta db error:", err.Error())
	}
	db, err = leveldb.OpenFile(dataPath, opts)
	if err != nil {
		log.Fatalln("open db error:", err.Error())
	}

	// start garbage collector routine.
	go gc()
}

// Close LevelDB
func Close() {
	db.Close()
}
