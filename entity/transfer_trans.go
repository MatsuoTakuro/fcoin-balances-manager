package entity

import (
	"database/sql"
	"time"
)

type TransferTransID uint64

type TransferTrans struct {
	ID          TransferTransID `db:"id"`
	FromUser    UserID          `db:"from_user"`    // TODO: 不要？Balanceの*structだけで良かったかも
	FromBalance BalanceID       `db:"from_balance"` // TODO: Balanceの*structのほうが良かった？かも
	ToUser      UserID          `db:"to_user"`      // TODO: 不要？Balanceの*structだけで良かったかも
	ToBalance   BalanceID       `db:"to_balance"`   // TODO: Balanceの*structのほうが良かった？かも
	Amount      uint32          `db:"amount"`
	ProcessedAt time.Time       `db:"processed_at"`
}

type TransferTransOpt struct {
	ID          sql.NullInt64 `db:"id"`
	FromUser    sql.NullInt64 `db:"from_user"`
	FromBalance sql.NullInt64 `db:"from_balance"`
	ToUser      sql.NullInt64 `db:"to_user"`
	ToBalance   sql.NullInt64 `db:"to_balance"`
	Amount      sql.NullInt32 `db:"amount"`
	ProcessedAt sql.NullTime  `db:"processed_at"`
}
