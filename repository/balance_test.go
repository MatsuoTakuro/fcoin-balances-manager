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

var balanceTable string = "balances"

func TestCreateBalance(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// トランザクションを張ることで、このテストケースの中だけのテーブル状態にする
	tx, err := utils.OpenDBForTest(t).BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	// このテストケースが完了したらもとに戻すため、ロールバックする
	t.Cleanup(func() { _ = tx.Rollback() })

	uname := randomUserName()
	fc := clock.FixedClocker{}
	wantBlance := &entity.Balance{
		UserID:    prepareUser(ctx, t, tx, uname, fc),
		Amount:    0,
		CreatedAt: fc.Now(),
		UpdatedAt: fc.Now(),
	}

	type mocks struct {
		puserID entity.UserID
	}

	tests := map[string]struct {
		wantBalance *entity.Balance
		wantErr     error
		mocks       mocks
	}{
		"normal": {
			wantBalance: wantBlance,
			wantErr:     nil,
			mocks: mocks{
				puserID: wantBlance.UserID,
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r := &Repository{Clocker: fc}
			gotBalance, gotErr := r.CreateBalance(ctx, tx, sub.mocks.puserID)

			if e := cmp.Diff(gotErr, sub.wantErr); len(e) != 0 {
				t.Errorf("differs: (-got +want)\n%s", e)
			}

			if b := cmp.Diff(gotBalance, sub.wantBalance,
				cmpopts.IgnoreFields(entity.Balance{}, "ID"),
			); len(b) != 0 {
				t.Errorf("differs: (-got +want)\n%s", b)
			}
		})
	}
}
