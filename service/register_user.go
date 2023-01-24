package service

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
)

type RegisterUserServiceImpl struct {
	DB   repository.Beginner
	Repo repository.UserRegisterRepo
}

func (ru *RegisterUserServiceImpl) RegisterUser(
	ctx context.Context, name string,
) (*entity.User, *entity.Balance, error) {

	tx, err := ru.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to begin trans for register_user: %w", err)
	}
	defer tx.Rollback()

	user, err := ru.Repo.CreateUser(ctx, tx, name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create user for register_user: %w", err)
	}

	balance, err := ru.Repo.CreateBalance(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create balance for register_user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, fmt.Errorf("failed to commit trans for register_user: %w", err)
	}

	return user, balance, nil
}
