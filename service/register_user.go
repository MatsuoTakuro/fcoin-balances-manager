package service

import (
	"context"

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

	user, balance, err := ru.Repo.UserRegisterWithTx(ctx, ru.DB, name)
	if err != nil {
		return nil, nil, err
	}
	return user, balance, err
}
