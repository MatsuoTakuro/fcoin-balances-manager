package entity

import (
	"time"
)

type TransferTransID uint64

type TransferTrans struct {
	ID          TransferTransID `db:"id"`
	FromUser    UserID          `db:"from_user"`    // TODO: 不要？Balanceの*structだけで良かったかも
	FromBalance BalanceID       `db:"from_balance"` // TODO: Balanceの*structのほうが良かった？かも（handler.respBodyも同様に変更？）
	ToUser      UserID          `db:"to_user"`      // TODO: 不要？Balanceの*structだけで良かったかも
	ToBalance   BalanceID       `db:"to_balance"`   // TODO: Balanceの*structのほうが良かった？かも（handler.respBodyも同様に変更？）
	Amount      uint32          `db:"amount"`
	ProcessedAt time.Time       `db:"processed_at"`
}
