package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

const (
	MySQLDuplicateEntryErrCode = 1062
)

var mysqlErr *mysql.MySQLError

func isDuplicateEntryErr(err error) bool {
	return errors.As(err, &mysqlErr) && mysqlErr.Number == MySQLDuplicateEntryErrCode
}

var noRowErr error = sql.ErrNoRows
