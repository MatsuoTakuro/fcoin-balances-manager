package repository

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/optional"
)

/*
コイン転送によらない残高更新トランザクションの作成
*/
func (r *Repository) CreateBalanceTrans(
	ctx context.Context, db Execer,
	userID entity.UserID, balanceID entity.BalanceID,
	amount int32,
) (*entity.BalanceTrans, error) {
	sql := `INSERT INTO balance_trans (
					user_id, balance_id, amount, processed_at
					) VALUES (?, ?, ?, ?)`

	processedAt := r.Clocker.Now()
	result, err := db.ExecContext(ctx, sql, userID, balanceID, amount, processedAt)
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to create balance_trans")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to get inserted balance_trans_id")
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

/*
コイン転送による残高更新トランザクションの作成
*/
func (r *Repository) CreateBalanceTransByTransfer(
	ctx context.Context, db Execer,
	userID entity.UserID, balanceID entity.BalanceID, TransferTransID entity.TransferTransID,
	amount int32,
) (*entity.BalanceTrans, error) {
	sql := `INSERT INTO balance_trans (
					user_id, balance_id, transfer_id, amount, processed_at
					) VALUES (?, ?, ?, ?, ?)`

	processedAt := r.Clocker.Now()
	result, err := db.ExecContext(ctx, sql, userID, balanceID, TransferTransID, amount, processedAt)
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to create balance_trans by tranfer")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to get inserted balance_trans_id")
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

func (r *Repository) GetBalanceTransListByBalanceID(
	ctx context.Context, db Queryer, balanceID entity.BalanceID,
) ([]*entity.BalanceTrans, error) {
	sql := `SELECT bt.id, bt.user_id, bt.balance_id,
						tt.id, tt.from_user, tt.from_balance, tt.to_user, tt.to_balance, tt.amount, tt.processed_at,
						bt.amount, bt.processed_at
					FROM balance_trans AS bt
					LEFT JOIN transfer_trans AS tt
					ON bt.transfer_id = tt.id
					WHERE bt.balance_id = ?
					ORDER BY bt.processed_at ASC`
	rows, err := db.QueryxContext(ctx, sql, balanceID)
	if err != nil {
		err = apperror.GET_DATA_FAILED.Wrap(err,
			fmt.Sprintf("failed to get balance_trans by balance_id: %d", balanceID))
		return nil, err
	}
	defer rows.Close()

	var btsOpt []*optional.BalanceTrans
	for rows.Next() {
		btOpt := optional.BalanceTrans{}
		if err := rows.Scan(
			&btOpt.ID,
			&btOpt.UserID,
			&btOpt.BalanceID,
			&btOpt.TransferTrans.ID,
			&btOpt.TransferTrans.FromUser,
			&btOpt.TransferTrans.FromBalance,
			&btOpt.TransferTrans.ToUser,
			&btOpt.TransferTrans.ToBalance,
			&btOpt.TransferTrans.Amount,
			&btOpt.TransferTrans.ProcessedAt,
			&btOpt.Amount,
			&btOpt.ProcessedAt,
		); err != nil {
			err = apperror.GET_DATA_FAILED.Wrap(err,
				fmt.Sprintf("failed to scan balance_trans gotten by balance_id: %d", balanceID))
			return nil, err
		}
		btsOpt = append(btsOpt, &btOpt)
	}

	var bts []*entity.BalanceTrans
	for _, btOpt := range btsOpt {
		bt := &entity.BalanceTrans{
			ID:          btOpt.ID,
			UserID:      btOpt.UserID,
			BalanceID:   btOpt.BalanceID,
			Amount:      btOpt.Amount,
			ProcessedAt: btOpt.ProcessedAt,
		}
		if !btOpt.TransferTrans.ID.Valid {
			bts = append(bts, bt)
			continue
		}
		bt.TransferTrans = entity.TransferTrans{
			ID:          entity.TransferTransID(uint64(btOpt.TransferTrans.ID.Int64)),
			FromUser:    entity.UserID(uint64(btOpt.TransferTrans.FromUser.Int64)),
			FromBalance: entity.BalanceID(uint64(btOpt.TransferTrans.FromBalance.Int64)),
			ToUser:      entity.UserID(uint64(btOpt.TransferTrans.ToUser.Int64)),
			ToBalance:   entity.BalanceID(uint64(btOpt.TransferTrans.ToBalance.Int64)),
			Amount:      uint32(btOpt.TransferTrans.Amount.Int32),
			ProcessedAt: btOpt.TransferTrans.ProcessedAt.Time,
		}
		bts = append(bts, bt)
	}

	return bts, nil
}
