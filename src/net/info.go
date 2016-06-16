package net

import (
	"strconv"

	"../redlot"
)

// Get server info.
func info(args [][]byte) (interface{}, error) {
	return "version:\n\t" + Version +
		"\nlinks:\n\t" + strconv.FormatUint(counter.ConnCounter, 10) +
		"\ncalls:\n\t" + strconv.FormatUint(counter.TotalCalls, 10) +
		"\n" + redlot.KeyRange(), nil
}
