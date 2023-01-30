package entity

import "time"

type BalanceTransID uint64

type BalanceTrans struct {
	ID            BalanceTransID `db:"id"`
	UserID        UserID         `db:"user_id"`
	BalanceID     BalanceID      `db:"balance_id"`
	TransferTrans TransferTrans  `db:"transfers"`
	Amount        int32          `db:"amount"`
	ProcessedAt   time.Time      `db:"processed_at"`
}
