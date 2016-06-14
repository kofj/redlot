package net

import "strconv"

// Get server info.
func Info(args [][]byte) (interface{}, error) {
	return "version:\n\t" + Version +
		"\nlinks:\n\t" + strconv.FormatUint(info.ConnCounter, 10) +
		"\ncalls:\n\t" + strconv.FormatUint(info.TotalCalls, 10) +
		"\n", nil
}
