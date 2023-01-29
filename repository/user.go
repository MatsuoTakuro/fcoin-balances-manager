package repository

import (
	"context"

	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
)

func (r *Repository) CreateUser(
	ctx context.Context, db Execer, name string,
) (*entity.User, error) {
	sql := `INSERT INTO users (
		name, created_at, updated_at
	) VALUES (?, ?, ?)`

	result, err := db.ExecContext(ctx, sql, name, r.Clocker.Now(), r.Clocker.Now())
	if err != nil {
		if isDuplicateEntryErr(err) {
			err = apperror.RegisterDuplicateDataRestricted.Wrap(err, "cannot create same name user")
			return nil, err
		}
		err = apperror.RegisterDataFailed.Wrap(err, "failed to create user")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = apperror.RegisterDataFailed.Wrap(err, "failed to get inserted user_id")
		return nil, err
	}

	user := &entity.User{
		ID:   entity.UserID(id),
		Name: name,
	}

	return user, nil
}
