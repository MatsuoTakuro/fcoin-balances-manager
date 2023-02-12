package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/clock"
)

type Repository struct {
	Clocker clock.Clocker
}

//go:generate go run github.com/matryer/moq -out repository_mock.go . UserRegisterRepo BalanceUpdaterRepo CoinsTransferRepo BalanceDetailsGetterRepo

type UserRegisterRepo interface {
	RegisterUserTx(ctx context.Context, db Beginner, name string) (*entity.User, *entity.Balance, error)
}

type BalanceUpdaterRepo interface {
	GetBalanceByUserID(ctx context.Context, db Queryer, userID entity.UserID) (*entity.Balance, error)
	UpdateBalanceTx(ctx context.Context, db Beginner, balance *entity.Balance, amount int32) (*entity.BalanceTrans, error)
}
type CoinsTransferRepo interface {
	GetBalanceByUserID(ctx context.Context, db Queryer, userID entity.UserID) (*entity.Balance, error)
	TransferCoinsTx(ctx context.Context, db Beginner, fromBalance *entity.Balance, toBalance *entity.Balance, amount uint32) (*entity.BalanceTrans, error)
}
type BalanceDetailsGetterRepo interface {
	GetBalanceByUserID(ctx context.Context, db Queryer, userID entity.UserID) (*entity.Balance, error)
	GetBalanceTransListByBalanceID(ctx context.Context, db Queryer, balanceID entity.BalanceID) ([]*entity.BalanceTrans, error)
}
