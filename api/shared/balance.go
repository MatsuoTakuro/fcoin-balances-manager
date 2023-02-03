package shared

import "github.com/MatsuoTakuro/fcoin-balances-manager/entity"

type Balance struct {
	ID     entity.BalanceID `json:"id,omitempty"`
	Amount uint32           `json:"amount"`
}
