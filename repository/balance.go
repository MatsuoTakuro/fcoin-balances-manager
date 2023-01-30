package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) CreateBalance(
	ctx context.Context, db Execer, userID entity.UserID,
) (*entity.Balance, error) {
	sql := `INSERT INTO balances (
		user_id, amount, created_at, updated_at
	) VALUES (?, ?, ?, ?)`

	result, err := db.ExecContext(ctx, sql, userID, 0, r.Clocker.Now(), r.Clocker.Now())
	if err != nil {
		if isDuplicateEntryErr(err) {
			err = apperror.RegisterDuplicateDataRestricted.Wrap(err, "can create only one balance per same user_id")
			return nil, err
		}
		err = apperror.RegisterDataFailed.Wrap(err, "failed to create balance")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to get inserted balance_id")
		return nil, err
	}

	balance := &entity.Balance{
		ID:     entity.BalanceID(id),
		UserID: userID,
		Amount: 0,
	}

	return balance, nil
}

func (r *Repository) GetBalanceByUserID(ctx context.Context, db Queryer, userID entity.UserID) (*entity.Balance, error) {
	return nil, nil
}

func (r *Repository) UpdateBalance(ctx context.Context, db Execer, balanceID entity.BalanceID, amount int32) error {
	return nil
}
