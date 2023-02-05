package entity

import "time"

type BalanceTransID uint64

type BalanceTrans struct {
	ID            BalanceTransID `db:"id"`
	UserID        UserID         `db:"user_id"`    // TODO: 不要？Balanceの*structだけで良かったかも
	BalanceID     BalanceID      `db:"balance_id"` // TODO: Balanceの*structのほうが良かった？かも（api.respBodyも同様に変更？）
	TransferTrans TransferTrans  `db:"transfers"`
	Amount        int32          `db:"amount"`
	ProcessedAt   time.Time      `db:"processed_at"`
}
