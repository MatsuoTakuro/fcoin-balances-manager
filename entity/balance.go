package entity

import "time"

type BalanceID uint64

type Balance struct {
	ID        BalanceID `db:"id"`
	UserID    UserID    `db:"user_id"`
	Amount    uint32    `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
