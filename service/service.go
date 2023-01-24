package service

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name string) (*entity.User, *entity.Balance, error)
}
