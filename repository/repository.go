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
	RegisterUserWithTx(ctx context.Context, db Beginner, name string) (*entity.User, *entity.Balance, error)
}

type BalanceUpdaterRepo interface {
	GetBalanceByUserID(ctx context.Context, db Queryer, userID entity.UserID) (*entity.Balance, error)
	UpdateBalanceWithTx(ctx context.Context, db Beginner, userID entity.UserID, balanceID entity.BalanceID, amount int32) (*entity.BalanceTrans, error)
}
