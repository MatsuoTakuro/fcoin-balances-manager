package entity

import "time"

type BalanceTransID uint64

type BalanceTrans struct {
	ID            BalanceTransID `db:"id"`
	UserID        UserID         `db:"user_id"`    // TODO: 不要？Balanceの*structだけで良かったかも
	BalanceID     BalanceID      `db:"balance_id"` // TODO: Balanceの*structのほうが良かった？かも
	TransferTrans TransferTrans  `db:"transfers"`
	Amount        int32          `db:"amount"`
	ProcessedAt   time.Time      `db:"processed_at"`
}
type BalanceTransOpt struct {
	ID            BalanceTransID   `db:"id"`
	UserID        UserID           `db:"user_id"`    // TODO: 不要？Balanceの*structだけで良かったかも
	BalanceID     BalanceID        `db:"balance_id"` // TODO: Balanceの*structのほうが良かった？かも
	TransferTrans TransferTransOpt `db:"transfers"`
	Amount        int32            `db:"amount"`
	ProcessedAt   time.Time        `db:"processed_at"`
}
