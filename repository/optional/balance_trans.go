package optional

import (
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

type BalanceTrans struct {
	ID            entity.BalanceTransID `db:"id"`
	UserID        entity.UserID         `db:"user_id"`    // TODO: 不要？Balanceの*structだけで良かったかも
	BalanceID     entity.BalanceID      `db:"balance_id"` // TODO: Balanceの*structのほうが良かった？かも（api.respBodyも同様に変更？）
	TransferTrans TransferTransOpt      `db:"transfers"`
	Amount        int32                 `db:"amount"`
	ProcessedAt   time.Time             `db:"processed_at"`
}
