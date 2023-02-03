package service

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
)

type GetBalanceDetailsServicer struct {
	DB   repository.Queryer
	Repo repository.BalanceDetailsGetterRepo
}

func (gb *GetBalanceDetailsServicer) GetBalanceDetails(
	ctx context.Context, userID entity.UserID,
) (*entity.Balance, []*entity.BalanceTrans, error) {

	balance, err := gb.Repo.GetBalanceByUserID(ctx, gb.DB, userID)
	if err != nil {
		return nil, nil, err
	}

	history, err := gb.Repo.GetBalanceTransListByBalanceID(ctx, gb.DB, balance.ID)
	if err != nil {
		return nil, nil, err
	}

	return balance, history, err
}
