package entity

import (
	"time"
)

type BalanceID uint64

type Balance struct {
	ID        BalanceID `db:"id"`
	UserID    UserID    `db:"user_id"` // TODO: Userの*structのほうが良かったかも（handler.respBodyも同様に変更？）
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

func (b *Balance) CanExceedMaxLimit(amount int32) bool {
	if amount > 0 {
		current := uint64(b.Amount)
		amount := uint64(amount)
		return (current + amount) > 2147483647
	}
	return false
}

func (b *Balance) UpdateAmount(amount int32) {
	if amount < 0 {
		b.Amount -= uint32(-amount)
	} else {
		b.Amount += uint32(amount)
	}
}
