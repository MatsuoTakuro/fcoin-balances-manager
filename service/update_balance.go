package service

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
)

type UpdateBalanceServicer struct {
	DB   repository.Beginner
	Repo repository.BalanceUpdaterRepo
}

func (ru *UpdateBalanceServicer) UpdateBalance(
	ctx context.Context, userID entity.UserID, amount int32,
) (*entity.BalanceTrans, error) {

	balance, err := ru.Repo.GetBalanceByUserID(ctx, ru.DB, userID)
	if err != nil {
		return nil, err
	}

	if !balance.CanBeZeroOrMore(amount) {
		err := apperror.NewAppError(apperror.AmountOverBalance,
			fmt.Sprintf("amount exceeds balance => input amount: %d, balance: %d", amount, balance.Amount))
		return nil, err
	}

	balanceTrans, err := ru.Repo.UpdateBalanceWithTx(ctx, ru.DB, userID, balance.ID, amount)
	if err != nil {
		return nil, err
	}

	return balanceTrans, err
}
