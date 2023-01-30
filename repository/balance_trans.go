package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) CreateBalanceTrans(
	ctx context.Context, db Execer, userID entity.UserID, balanceID entity.BalanceID, amount int32,
) (*entity.BalanceTrans, error) {
	sql := `INSERT INTO balance_trans (
					user_id, balance_id, amount, processed_at
					) VALUES (?, ?, ?, ?)`

	processedAt := r.Clocker.Now()
	result, err := db.ExecContext(ctx, sql, userID, balanceID, amount, processedAt)
	if err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to create balance_trans")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to get inserted balance_trans_id")
		return nil, err
	}

	balanceTrans := &entity.BalanceTrans{
		ID:          entity.BalanceTransID(id),
		UserID:      userID,
		BalanceID:   balanceID,
		Amount:      amount,
		ProcessedAt: processedAt,
	}

	return balanceTrans, nil
}
