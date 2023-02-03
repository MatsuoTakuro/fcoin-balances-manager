package shared

import (
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

type BalanceTrans struct {
	ID          entity.BalanceTransID `json:"id,omitempty"`
	UserID      entity.UserID         `json:"user_id,omitempty"`
	BalanceID   entity.BalanceID      `json:"balance_id,omitempty"`
	Transfer    *TransferTrans        `json:"transfer,omitempty"`
	Amount      int32                 `json:"amount,omitempty"`
	ProcessedAt time.Time             `json:"processed_at,omitempty"`
}
