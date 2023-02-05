package repository

/*
更新系のトランザクション処理についてまとめる
*/

import (
	"context"
	"database/sql"
	"log"

	"github.com/MatsuoTakuro/fcoin-balances-manager/appcontext"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) BeginTx(ctx context.Context, db Beginner) (*sql.Tx, func(context.Context, *sql.Tx, error) error, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, err.Error())
	}
	return tx, commitOrRollback, nil
}

func commitOrRollback(ctx context.Context, tx *sql.Tx, err error) error {
	if err == nil {
		if err := tx.Commit(); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, rbErr.Error())
			}
			err = apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, err.Error())
			log.Printf("[%d]trans: rollbacked sucessfully for handling err: %v",
				appcontext.GetTracdID(ctx), err)
			return err
		}
		return nil
	} else {
		if rbErr := tx.Rollback(); rbErr != nil {
			return apperror.PROCESS_TRANSACTION_FAILED.Wrap(err, rbErr.Error())
		}
		log.Printf("[%d]trans: rollbacked sucessfully for handling err: %v",
			appcontext.GetTracdID(ctx), err)
		return err
	}
}

func (r *Repository) RegisterUserTx(
	ctx context.Context, db Beginner, name string,
) (user *entity.User, balance *entity.Balance, err error) {

	tx, commitOrRollback, err := r.BeginTx(ctx, db)
	defer func() {
		if txErr := commitOrRollback(ctx, tx, err); txErr != nil {
			err = txErr
		}
	}()

	user, err = r.CreateUser(ctx, tx, name)
	if err != nil {
		return nil, nil, err
	}

	balance, err = r.CreateBalance(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, balance, nil
}

// （コイン転送によらない）残高更新トランザクション処理
func (r *Repository) UpdateBalanceTx(
	ctx context.Context, db Beginner, balance *entity.Balance, amount int32,
) (balanceTrans *entity.BalanceTrans, err error) {

	tx, commitOrRollback, err := r.BeginTx(ctx, db)
	defer func() {
		if txErr := commitOrRollback(ctx, tx, err); txErr != nil {
			err = txErr
		}
	}()

	balanceTrans, err = r.CreateBalanceTrans(ctx, tx, balance.UserID, balance.ID, amount)
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(ctx, tx, balance.ID, balance.Amount)
	if err != nil {
		return nil, err
	}

	return balanceTrans, nil
}

func (r *Repository) TransferCoinsTx(
	ctx context.Context, db Beginner, fromBalance *entity.Balance, toBalance *entity.Balance, amount uint32,
) (*entity.BalanceTrans, error) {

	tx, commitOrRollback, err := r.BeginTx(ctx, db)
	defer func() {
		if txErr := commitOrRollback(ctx, tx, err); txErr != nil {
			err = txErr
		}
	}()

	transferTrans, err := r.CreateTransferTrans(
		ctx, tx, fromBalance, toBalance, amount)
	if err != nil {
		return nil, err
	}

	balanceTrans, err := r.CreateBalanceTransByTransfer(
		ctx, tx, transferTrans.FromUser, transferTrans.FromBalance, transferTrans.ID, -int32(amount))
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(
		ctx, tx, fromBalance.ID, fromBalance.Amount)
	if err != nil {
		return nil, err
	}

	_, err = r.CreateBalanceTransByTransfer(
		ctx, tx, transferTrans.ToUser, transferTrans.ToBalance, transferTrans.ID, int32(amount))
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(
		ctx, tx, toBalance.ID, toBalance.Amount)
	if err != nil {
		return nil, err
	}

	balanceTrans.TransferTrans = *transferTrans

	return balanceTrans, nil
}
