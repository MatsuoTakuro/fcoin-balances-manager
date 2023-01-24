package entity

import "time"

type BalanceID uint64

type Balance struct {
	ID        BalanceID `json:"id" db:"id"`
	UserID    UserID    `json:"user_id" db:"user_id"`
	Amount    uint32    `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
