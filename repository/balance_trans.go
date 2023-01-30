package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) CreateBalanceTrans(
	ctx context.Context, tx Execer, userID entity.UserID, balanceID entity.BalanceID,
) (*entity.BalanceTrans, error) {
	return nil, nil
}
