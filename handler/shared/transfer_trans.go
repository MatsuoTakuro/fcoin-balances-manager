package shared

import (
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

type TransferTrans struct {
	ID          entity.TransferTransID `json:"id,omitempty"`
	FromUser    entity.UserID          `json:"from_user,omitempty"`
	FromBalance entity.BalanceID       `json:"from_balance,omitempty"`
	ToUser      entity.UserID          `json:"to_user,omitempty"`
	ToBalance   entity.BalanceID       `json:"to_balance,omitempty"`
	Amount      uint32                 `json:"amount,omitempty"`
	ProcessedAt time.Time              `json:"processed_at,omitempty"`
}
