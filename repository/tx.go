package repository

import (
	"context"
	"log"

	"github.com/MatsuoTakuro/fcoin-balances-manager/appcontext"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) RegisterUserWithTx(
	ctx context.Context, db Beginner, name string,
) (user *entity.User, balance *entity.Balance, err error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to begin transaction for registering user")
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				rbErr = apperror.RegisterDataFailed.Wrap(err, err.Error())
				err = apperror.RegisterDataFailed.Wrap(rbErr, "failed to rollback transaction for registering user")
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

	if err = tx.Commit(); err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to commit transaction for updating balance")
		return nil, nil, err
	}

	return user, balance, nil
}

func (r *Repository) UpdateBalanceWithTx(
	ctx context.Context, db Beginner, userID entity.UserID, balanceID entity.BalanceID, amount int32, balanceAmount uint32,
) (balanceTrans *entity.BalanceTrans, err error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = apperror.UpdateDataFailed.Wrap(err, "failed to begin transaction for updating balance")
		return nil, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				rbErr = apperror.UpdateDataFailed.Wrap(err, err.Error())
				err = apperror.UpdateDataFailed.Wrap(rbErr, "failed to rollback transaction for updating balance")
				return
			}
			log.Printf("[%d]trans: rollbacked sucessfully for updating balance", appcontext.GetTracdID(ctx))
		}
	}()

	balanceTrans, err = r.CreateBalanceTransWithoutTransfer(ctx, tx, userID, balanceID, amount)
	if err != nil {
		return nil, err
	}

	err = r.UpdateBalanceByID(ctx, tx, balanceID, balanceAmount)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		err = apperror.UpdateDataFailed.Wrap(err, "failed to commit transaction for updating balance")
		return nil, err
	}

	return balanceTrans, nil
}
