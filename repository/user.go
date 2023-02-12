package repository

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) CreateUser(
	ctx context.Context, db Execer, name string,
) (*entity.User, error) {
	sql := `INSERT INTO users (
		name, created_at, updated_at
	) VALUES (?, ?, ?)`

	currentTime := r.Clocker.Now()
	result, err := db.ExecContext(ctx, sql, name, currentTime, currentTime)
	if err != nil {
		if isDuplicateEntryErr(err) {
			return nil, apperror.REGISTER_DUPLICATE_DATA_RESTRICTED.Wrap(err, fmt.Sprintf("cannot create same name user: %s", name))
		}
		return nil, apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to create user")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, apperror.REGISTER_DATA_FAILED.Wrap(err, "failed to get inserted user_id")
	}

	user := &entity.User{
		ID:        entity.UserID(id),
		Name:      name,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	return user, nil
}
