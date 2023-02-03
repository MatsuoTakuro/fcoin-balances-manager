package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) CreateTransferTrans(
	ctx context.Context, db Execer, fromBalance *entity.Balance, toBalance *entity.Balance, amount uint32,
) (*entity.TransferTrans, error) {
	sql := `INSERT INTO transfer_trans (
		from_user, from_balance, to_user, to_balance, amount, processed_at
		) VALUES (?, ?, ?, ?, ?, ?)`

	processedAt := r.Clocker.Now()
	result, err := db.ExecContext(ctx, sql, fromBalance.UserID, fromBalance.ID, toBalance.UserID, toBalance.ID, amount, processedAt)
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to create transfer_trans")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to get inserted transfer_trans_id")
		return nil, err
	}

	transferTrans := &entity.TransferTrans{
		ID:          entity.TransferTransID(id),
		FromUser:    fromBalance.UserID,
		FromBalance: fromBalance.ID,
		ToUser:      toBalance.UserID,
		ToBalance:   toBalance.ID,
		Amount:      amount,
		ProcessedAt: processedAt,
	}

	return transferTrans, nil
}
