package repository

import (
	"context"
	"errors"
	"fmt"

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
			err = apperror.RegisterDuplicateDataRestricted.Wrap(err,
				fmt.Sprintf("can create only one balance per same user_id: %d", userID))
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
	sql := `SELECT id, amount FROM balances
					WHERE user_id = ?`
	result := db.QueryRowxContext(ctx, sql, userID)
	balance := &entity.Balance{}
	err := result.Scan(&balance.ID, &balance.Amount)
	if err != nil {
		if errors.Is(err, noRowErr) {
			err = apperror.NoSelectedData.Wrap(err,
				fmt.Sprintf("no selected balance by user_id: %d", userID))
			return nil, err
		} else {
			err = apperror.GetDataFailed.Wrap(err,
				fmt.Sprintf("failed to get balance by user_id: %d", userID))
			return nil, err
		}
	}

	return balance, nil
}

func (r *Repository) UpdateBalanceByID(ctx context.Context, db Execer, balanceID entity.BalanceID, amount uint32) error {
	sql := `UPDATE balances SET
					amount = ?, updated_at = ?
					WHERE id = ?`

	_, err := db.ExecContext(ctx, sql, amount, r.Clocker.Now(), balanceID)
	if err != nil {
		err = apperror.UpdateDataFailed.Wrap(err,
			fmt.Sprintf("failed to update balance by id: %d", balanceID))
		return err
	}
	return nil
}
