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

	currentTime := r.Clocker.Now()
	result, err := db.ExecContext(ctx, sql, userID, 0, currentTime, currentTime)
	if err != nil {
		if isDuplicateEntryErr(err) {
			return nil, apperror.REGISTER_DUPLICATE_DATA_RESTRICTED.Wrap(err, fmt.Sprintf("can create only one balance per same user_id: %d", userID))
		}
		return nil, apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to create balance")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to get inserted balance_id")
	}

	balance := &entity.Balance{
		ID:        entity.BalanceID(id),
		UserID:    userID,
		Amount:    0,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	return balance, nil
}

func (r *Repository) GetBalanceByUserID(ctx context.Context, db Queryer, userID entity.UserID) (*entity.Balance, error) {
	sql := `SELECT id, amount FROM balances
					WHERE user_id = ?`
	result := db.QueryRowxContext(ctx, sql, userID)
	balance := &entity.Balance{
		UserID: userID,
	}
	err := result.Scan(&balance.ID, &balance.Amount)
	if err != nil {
		if errors.Is(err, noRowErr) {
			return nil, apperror.NO_SELECTED_DATA.Wrap(err, fmt.Sprintf("not found balance selected by user_id: %d", userID))
		} else {
			return nil, apperror.GET_DATA_FAILED.Wrap(err, fmt.Sprintf("failed to get balance by user_id: %d", userID))
		}
	}

	return balance, nil
}

func (r *Repository) UpdateBalanceByID(ctx context.Context, db Execer, balanceID entity.BalanceID, amount uint32) error {
	sql := `UPDATE balances SET
					amount = ?, updated_at = ?
					WHERE id = ?`

	// TODO: 更新日時や金額等で事前に更新が無いかを確認する
	_, err := db.ExecContext(ctx, sql, amount, r.Clocker.Now(), balanceID)
	if err != nil {
		return apperror.UPDATE_DATA_FAILED.Wrap(err, fmt.Sprintf("failed to update balance by id: %d", balanceID))
	}
	return nil
}
