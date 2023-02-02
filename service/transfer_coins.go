package service

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
)

type TransferCoinsServicer struct {
	DB   repository.Beginner
	Repo repository.TransferCoinsRepo
}

func (tc *TransferCoinsServicer) TransferCoins(
	ctx context.Context, fromUser entity.UserID, toUser entity.UserID, amount uint32,
) (*entity.BalanceTrans, error) {

	if fromUser == toUser {
		err := apperror.NewAppError(apperror.NoTransferOfCoinsBySameUser,
			fmt.Sprintf(
				"users are the same source and destination of transferring coins => input user_id as source: %d, user_id as destination: %d",
				fromUser, toUser))
		return nil, err
	}

	fromBalance, err := tc.Repo.GetBalanceByUserID(ctx, tc.DB, fromUser)
	if err != nil {
		return nil, err
	}

	if !fromBalance.CanBeZeroOrMore(-int32(amount)) {
		err := apperror.NewAppError(apperror.ConsumedAmountOverBalance,
			fmt.Sprintf("amount consumed exceeds current balance => input amount: %d, balance: %d", amount, fromBalance.Amount))
		return nil, err
	}
	fromBalance.UpdateAmount(-int32(amount))

	toBalance, err := tc.Repo.GetBalanceByUserID(ctx, tc.DB, toUser)
	if err != nil {
		return nil, err
	}
	if toBalance.CanExceedMaxLimit(int32(amount)) {
		err := apperror.NewAppError(apperror.OverMaxBalanceLimit,
			fmt.Sprintf("total amount exceeds max balance limit => input amount: %d, transfered to user_id: %d", amount, toBalance.UserID))
		return nil, err
	}
	toBalance.UpdateAmount(int32(amount))

	balanceTrans, err := tc.Repo.TransferCoinsTx(ctx, tc.DB, fromBalance, toBalance, amount)
	if err != nil {
		return nil, err
	}

	return balanceTrans, err
}
