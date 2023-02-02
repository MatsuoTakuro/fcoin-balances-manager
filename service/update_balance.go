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
		err := apperror.NewAppError(apperror.ConsumedAmountOverBalance,
			fmt.Sprintf("amount consumed exceeds current balance => input amount: %d, balance: %d", amount, balance.Amount))
		return nil, err
	}
	if balance.CanExceedMaxLimit(amount) {
		err := apperror.NewAppError(apperror.OverMaxBalanceLimit,
			fmt.Sprintf("total amount exceeds max balance limit => input amount: %d, balance: %d", amount, balance.Amount))
		return nil, err
	}
	balance.UpdateAmount(amount)

	balanceTrans, err := ru.Repo.UpdateBalanceTx(ctx, ru.DB, balance, amount)
	if err != nil {
		return nil, err
	}

	return balanceTrans, err
}
