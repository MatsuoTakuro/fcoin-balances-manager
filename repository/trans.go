package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) UserRegisterWithTx(
	ctx context.Context, db Beginner, name string,
) (*entity.User, *entity.Balance, error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to begin transaction for registering user")
		return nil, nil, err
	}
	// TODO: deferのRollbackエラーを拾う方法を調べる
	defer func() error {
		if err := tx.Rollback(); err == nil {
			err = apperror.RegisterDataFailed.Wrap(err, "failed to rollback transaction for registering user")
			return err
		}
		return nil
	}()

	user, err := r.CreateUser(ctx, tx, name)
	if err != nil {
		return nil, nil, err
	}

	balance, err := r.CreateBalance(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to commit transaction for registering user")
		return nil, nil, err
	}

	return user, balance, nil
}
