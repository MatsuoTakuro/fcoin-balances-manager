package repository

import "errors"

const (
	MySQLDuplicateEntryErrCode = 1062
)

var (
	ErrAlreadyEntry = errors.New("duplicate entry error")
)
