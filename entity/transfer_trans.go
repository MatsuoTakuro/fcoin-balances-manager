package entity

import "time"

type TransferTransID uint64

type TransferTrans struct {
	ID          TransferTransID `db:"id"`
	FromUser    UserID          `db:"from_user"`    // TODO: 不要？Balanceのstructだけで良かったかも
	FromBalance BalanceID       `db:"from_balance"` // TODO: Balanceのstructのほうが良かった？かも
	ToUser      UserID          `db:"to_user"`      // TODO: 不要？Balanceのstructだけで良かったかも
	ToBalance   BalanceID       `db:"to_balance"`   // TODO: Balanceのstructのほうが良かった？かも
	Amount      uint32          `db:"amount"`
	ProcessedAt time.Time       `db:"processed_at"`
}
