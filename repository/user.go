package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/go-sql-driver/mysql"
)

func (r *Repository) CreateUser(
	ctx context.Context, db Execer, name string,
) (*entity.User, error) {
	sql := `INSERT INTO users (
		name, created_at, updated_at
	) VALUES (?, ?, ?)`

	result, err := db.ExecContext(ctx, sql, name, r.Clocker.Now(), r.Clocker.Now())
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == MySQLDuplicateEntryErrCode {
			return nil, fmt.Errorf("cannot create same name user: %w", ErrAlreadyEntry)
		}
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:   entity.UserID(id),
		Name: name,
	}

	return user, nil
}
