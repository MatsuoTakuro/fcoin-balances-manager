package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/clock"
)

type Repository struct {
	Clocker clock.Clocker
}

type UserRegisterRepo interface {
	UserRegisterWithTx(ctx context.Context, db Beginner, name string) (*entity.User, *entity.Balance, error)
}
