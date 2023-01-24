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
	CreateUser(ctx context.Context, db Execer, name string) (*entity.User, error)
	CreateBalance(ctx context.Context, db Execer, userID entity.UserID) (*entity.Balance, error)
}
