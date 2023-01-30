package service

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
)

type RegisterUserServicer struct {
	DB   repository.Beginner
	Repo repository.UserRegisterRepo
}

func (ru *RegisterUserServicer) RegisterUser(
	ctx context.Context, name string,
) (*entity.User, *entity.Balance, error) {

	user, balance, err := ru.Repo.RegisterUserWithTx(ctx, ru.DB, name)
	if err != nil {
		return nil, nil, err
	}
	return user, balance, err
}
