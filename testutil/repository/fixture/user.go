package fixture

import (
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func User(u *entity.User) *entity.User {
	result := &entity.User{}

	if u == nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if !u.CreatedAt.IsZero() {
		result.CreatedAt = u.CreatedAt
	}
	if !u.UpdatedAt.IsZero() {
		result.UpdatedAt = u.UpdatedAt
	}
	return result
}
