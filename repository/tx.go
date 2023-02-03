package repository

/*
更新系のトランザクション処理についてまとめる
*/

import (
	"context"
	"log"

	"github.com/MatsuoTakuro/fcoin-balances-manager/appcontext"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) RegisterUserTx(
	ctx context.Context, db Beginner, name string,
) (user *entity.User, balance *entity.Balance, err error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to begin transaction for registering user")
		return nil, nil, err
	}
	defer func() {
		if cErr := tx.Commit(); cErr != nil {
			cErr = apperror.UPDATE_DATA_FAILED.Wrap(err, err.Error())
			err = apperror.UPDATE_DATA_FAILED.Wrap(cErr, "failed to commit transaction for registering user")
		}

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				rbErr = apperror.REGISTER_DATA_FAILED.Wrap(err, err.Error())
				err = apperror.REGISTER_DATA_FAILED.Wrap(rbErr, "failed to rollback transaction for registering user")
				return
			}
			log.Printf("[%d]trans: rollbacked sucessfully for registering user", appcontext.GetTracdID(ctx))
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

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = apperror.UPDATE_DATA_FAILED.Wrap(err, "failed to begin transaction for updating balance")
		return nil, err
	}
	defer func() {
		if cErr := tx.Commit(); cErr != nil {
			cErr = apperror.UPDATE_DATA_FAILED.Wrap(err, err.Error())
			err = apperror.UPDATE_DATA_FAILED.Wrap(cErr, "failed to commit transaction for updating balance")
		}

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				rbErr = apperror.UPDATE_DATA_FAILED.Wrap(err, err.Error())
				err = apperror.UPDATE_DATA_FAILED.Wrap(rbErr, "failed to rollback transaction for updating balance")
				return
			}
			log.Printf("[%d]trans: rollbacked sucessfully for updating balance", appcontext.GetTracdID(ctx))
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

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = apperror.UPDATE_DATA_FAILED.Wrap(err, "failed to begin transaction for transferring coins")
		return nil, err
	}
	defer func() {
		if cErr := tx.Commit(); cErr != nil {
			cErr = apperror.UPDATE_DATA_FAILED.Wrap(err, err.Error())
			err = apperror.UPDATE_DATA_FAILED.Wrap(cErr, "failed to commit transaction for transferring coins")
		}

		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				rbErr = apperror.UPDATE_DATA_FAILED.Wrap(err, err.Error())
				err = apperror.UPDATE_DATA_FAILED.Wrap(rbErr, "failed to rollback transaction for transferring coins")
				return
			}
			log.Printf("[%d]trans: rollbacked sucessfully for transfer coins", appcontext.GetTracdID(ctx))
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
