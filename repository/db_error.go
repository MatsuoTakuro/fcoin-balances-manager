package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DUPLICATE_ENTRY_ERRCODE = 1062
)

var mysqlErr *mysql.MySQLError

func isDuplicateEntryErr(err error) bool {
	return errors.As(err, &mysqlErr) && mysqlErr.Number == MYSQL_DUPLICATE_ENTRY_ERRCODE
}

var noRowErr error = sql.ErrNoRows
