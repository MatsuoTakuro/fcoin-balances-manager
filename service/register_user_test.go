package service

import (
	"context"
	"errors"
	"testing"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
	"github.com/google/go-cmp/cmp"
)

func TestRegisterUserServicer_RegisterUser(t *testing.T) {
	t.Parallel()

	wantUID := entity.UserID(1)
	wantName := "taro"
	type result struct {
		user    *entity.User
		balance *entity.Balance
		err     error
	}
	want := &result{
		user: &entity.User{
			ID:   wantUID,
			Name: wantName,
		},
		balance: &entity.Balance{
			UserID: wantUID,
			Amount: 0,
		},
		err: nil,
	}

	tests := map[string]struct {
		user    *entity.User
		balance *entity.Balance
		others  RegisterUserOthers
	}{
		"ok": {
			others: RegisterUserOthers{
				inputName:   wantName,
				expectedErr: nil,
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dbMock := &repository.BeginnerMock{}
			repoMock := &repository.UserRegisterRepoMock{}
			repoMock.RegisterUserTxFunc = func(
				ctx context.Context,
				db repository.Beginner,
				name string,
			) (*entity.User, *entity.Balance, error) {

				if db != dbMock {
					t.Fatalf("not want db %v", db)
				}
				if d := cmp.Diff(name, sub.others.inputName); len(d) != 0 {
					t.Fatalf("differs: (-got +want)\n%s", d)
				}

				if sub.others.expectedErr == nil {
					return &entity.User{
							ID:   1,
							Name: "taro",
						}, &entity.Balance{
							UserID: 1,
							Amount: 0,
						}, nil
				}

				return nil, nil, errors.New("error from UserRegisterRepoMock")
			}

			ru := &RegisterUserServicer{
				DB:   dbMock,
				Repo: repoMock,
			}

			gotUser, gotBalance, err := ru.RegisterUser(context.Background(), sub.want.Name)
			if err != nil {
				t.Fatalf("want no error, but got %v", err)
				return
			}
			if d := cmp.Diff(got, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}

		})
	}
}
