package service

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

//go:generate go run github.com/matryer/moq -out service_mock.go . RegisterUserService UpdateBalanceService TransferCoinsService GetBalanceDetails

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name string) (*entity.User, *entity.Balance, error)
}
type UpdateBalanceService interface {
	UpdateBalance(ctx context.Context, userID entity.UserID, amount int32) (*entity.BalanceTrans, error)
}
type TransferCoinsService interface {
	TransferCoins(ctx context.Context, fromUser entity.UserID, toUser entity.UserID, amount uint32) (*entity.BalanceTrans, error)
}
type GetBalanceDetails interface {
	GetBalanceDetails(ctx context.Context, userID entity.UserID) (*entity.Balance, []*entity.BalanceTrans, error)
}
