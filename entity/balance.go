package entity

import (
	"time"
)

type BalanceID uint64

type Balance struct {
	ID        BalanceID `db:"id"`
	UserID    UserID    `db:"user_id"`
	Amount    uint32    `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 残高が（プラスおよびマイナスの）更新後に、0以上になるかをチェックする
func (b *Balance) CanBeZeroOrMore(amount int32) bool {
	if amount < 0 {
		return b.Amount >= uint32(-amount)
	}
	return true
}
