package net

import "strconv"

func Info(args [][]byte) ([][]byte, error) {
	info := "version:\n\t" + Version +
		"\nlinks:\n\t" + strconv.FormatUint(info.ConnCounter, 10) +
		"\ncalls:\n\t" + strconv.FormatUint(info.TotalCalls, 10) +
		"\n"
	return [][]byte{[]byte(info)}, nil
}
