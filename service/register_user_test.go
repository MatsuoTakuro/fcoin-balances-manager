package service

import (
	"context"
	"testing"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository"
	"github.com/google/go-cmp/cmp"
)

func TestRegisterUser(t *testing.T) {
	t.Parallel()

	wantUser := &entity.User{
		ID:   1,
		Name: "taro",
	}
	wantBalance := &entity.Balance{
		UserID: 1,
		Amount: 0,
	}
	type otherMocks struct {
		name string
		err  error
	}

	tests := map[string]struct {
		wantUser    *entity.User
		wantBalance *entity.Balance
		wantErr     error
		others      otherMocks
	}{
		"normal": {
			wantUser:    wantUser,
			wantBalance: wantBalance,
			wantErr:     nil,
			others: otherMocks{
				name: "taro",
				err:  nil,
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			mockDb := &repository.BeginnerMock{}
			mockRepo := &repository.UserRegisterRepoMock{}
			mockRepo.RegisterUserTxFunc = func(
				pctx context.Context, pdb repository.Beginner, pname string,
			) (*entity.User, *entity.Balance, error) {
				if ctx != pctx {
					t.Fatalf("want context: %v ,got %v", ctx, pctx)
				}
				if mockDb != pdb {
					t.Fatalf("want db: %v ,got %v", mockDb, pdb)
				}
				if sub.others.name != pname {
					t.Fatalf("want name: %s ,got %s", sub.others.name, pname)
				}
				return &entity.User{
						ID:   1,
						Name: "taro",
					}, &entity.Balance{
						UserID: 1,
						Amount: 0,
					}, nil
			}

			ru := &RegisterUserServicer{
				DB:   mockDb,
				Repo: mockRepo,
			}
			gotUser, gotBalance, gotErr := ru.RegisterUser(ctx, sub.others.name)
			if gotErr != sub.others.err {
				t.Fatalf("want error(including nil): %v, but got %v", sub.others.err, gotErr)
			}
			if u := cmp.Diff(gotUser, sub.wantUser); len(u) != 0 {
				t.Errorf("differs: (-got +want)\n%s", u)
			}
			if b := cmp.Diff(gotBalance, sub.wantBalance); len(b) != 0 {
				t.Errorf("differs: (-got +want)\n%s", b)
			}
		})
	}
}
