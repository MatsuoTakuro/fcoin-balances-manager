package entity

import "time"

type TransferTransID uint64

type TransferTrans struct {
	ID          TransferTransID `db:"id"`
	FromUser    UserID          `db:"from_user"`
	FromBalance BalanceID       `db:"from_balance"`
	ToUser      UserID          `db:"to_user"`
	ToBalance   BalanceID       `db:"to_balance"`
	Amount      uint32          `db:"amount"`
	ProcessedAt time.Time       `db:"processed_at"`
}
