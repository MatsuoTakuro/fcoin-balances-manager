package repository

import (
	"context"
	"testing"

	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/clock"
	utils "github.com/MatsuoTakuro/fcoin-balances-manager/testutil/repository"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRegisterUserTx(t *testing.T) {
	t.Parallel()

	// TODO: テーブル状態に影響が無いように、トランザクション処理のテストをしたい
	db := utils.OpenDBForTest(t)
	t.Cleanup(func() {
		ctx := context.Background()
		tables := []string{balanceTable, userTable}
		err := utils.TruncateTables(ctx, db, tables)
		if err != nil {
			t.Fatal(err)
		}
	})

	fc := clock.FixedClocker{}

	normalUser := &entity.User{
		Name:      randomUserName(),
		CreatedAt: fc.Now(),
		UpdatedAt: fc.Now(),
	}
	normalBalance := &entity.Balance{
		Amount:    0,
		CreatedAt: fc.Now(),
		UpdatedAt: fc.Now(),
	}
	type mocks struct {
		pname string
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
				pname: normalUser.Name,
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			r := &Repository{Clocker: fc}

			gotUser, gotBalance, gotErr := r.RegisterUserTx(ctx, db, sub.mocks.pname)
			if gotErr != sub.wantErr {
				t.Errorf("want error(including nil): %v, but got %v", sub.wantErr, gotErr)
			}

			if u := cmp.Diff(gotUser, sub.wantUser,
				cmpopts.IgnoreFields(entity.User{}, "ID"),
			); len(u) != 0 {
				t.Errorf("differs: (-got +want)\n%s", u)
			}

			if b := cmp.Diff(gotBalance, sub.wantBalance,
				cmpopts.IgnoreFields(entity.Balance{}, "ID"),
				cmpopts.IgnoreFields(entity.Balance{}, "UserID"),
			); len(b) != 0 {
				t.Errorf("differs: (-got +want)\n%s", b)
			}

			if gotUser.ID != gotBalance.UserID {
				t.Errorf("want same user_id: %d in user, but %d in balance", gotUser.ID, gotBalance.UserID)
			}
		})
	}
}
