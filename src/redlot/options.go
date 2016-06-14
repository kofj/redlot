package redlot

import (
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Options struct {
	DataPath          string
	CacheSize         int
	BlockSize         int
	WriteBuffer       int
	CompactionBackoff bool
}

// convert redlot options to goleveldb options.
func (c *Options) convert() *opt.Options {
	return &opt.Options{
		BlockCacheCapacity:       c.CacheSize * opt.MiB,
		BlockSize:                c.BlockSize * opt.KiB,
		DisableCompactionBackoff: !c.CompactionBackoff,
		Filter:      filter.NewBloomFilter(10),
		WriteBuffer: c.WriteBuffer * opt.MiB,
	}
}
