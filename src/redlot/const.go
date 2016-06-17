package redlot

import "errors"

const (
	typeKV    = 'k'
	typeHASH  = 'h'
	typeHSIZE = 'H'
)

var (
	errNosArgs = errors.New("wrong number of arguments")
	errNotInt  = errors.New("value is not an integer or out of range")
)
