package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/go-sql-driver/mysql"
)

func (r *Repository) CreateBalance(
	ctx context.Context, db Execer, userID entity.UserID,
) (*entity.Balance, error) {
	sql := `INSERT INTO balances (
		user_id, amount, created_at, updated_at
	) VALUES (?, ?, ?, ?)`

	result, err := db.ExecContext(ctx, sql, userID, 0, r.Clocker.Now(), r.Clocker.Now())
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == MySQLDuplicateEntryErrCode {
			return nil, fmt.Errorf("can create only one balance per same user_id: %w", ErrAlreadyEntry)
		}
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	balance := &entity.Balance{
		ID:     entity.BalanceID(id),
		UserID: userID,
		Amount: 0,
	}

	return balance, nil
}
