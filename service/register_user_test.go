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

	normalUser := &entity.User{
		ID:   1,
		Name: "taro",
	}
	normalBalance := &entity.Balance{
		UserID: normalUser.ID,
		Amount: 0,
	}
	type mocks struct {
		pname    string
		ruser    *entity.User
		rbalance *entity.Balance
		rerr     error
	}

	tests := map[string]struct {
		wantUser    *entity.User
		wantBalance *entity.Balance
		wantErr     error
		mocks       mocks
	}{
		"normal": {
			wantUser:    normalUser,
			wantBalance: normalBalance,
			wantErr:     nil,
			mocks: mocks{
				pname:    normalUser.Name,
				ruser:    normalUser,
				rbalance: normalBalance,
				rerr:     nil,
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
				if sub.mocks.pname != pname {
					t.Fatalf("want name: %s ,got %s", sub.mocks.pname, pname)
				}
				return sub.mocks.ruser, sub.mocks.rbalance, sub.mocks.rerr
			}

			ru := &RegisterUserServicer{
				DB:   mockDb,
				Repo: mockRepo,
			}
			gotUser, gotBalance, gotErr := ru.RegisterUser(ctx, sub.mocks.pname)
			if gotErr != sub.mocks.rerr {
				t.Fatalf("want error(including nil): %v, but got %v", sub.mocks.rerr, gotErr)
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
