package repository

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) UserRegisterWithTx(
	ctx context.Context, db Beginner, name string,
) (*entity.User, *entity.Balance, error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to begin trans for register_user: %w", err)
	}
	defer tx.Rollback()

	user, err := r.CreateUser(ctx, tx, name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create user for register_user: %w", err)
	}

	balance, err := r.CreateBalance(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create balance for register_user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, fmt.Errorf("failed to commit trans for register_user: %w", err)
	}

	return user, balance, nil
}
