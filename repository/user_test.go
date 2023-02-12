package repository

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MatsuoTakuro/fcoin-balances-manager/apperror"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"github.com/MatsuoTakuro/fcoin-balances-manager/repository/clock"
	utils "github.com/MatsuoTakuro/fcoin-balances-manager/testutil/repository"
	"github.com/MatsuoTakuro/fcoin-balances-manager/testutil/repository/fixture"
	"github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jmoiron/sqlx"
)

var userTable string = "users"

func TestCreateUser(t *testing.T) {
	t.Parallel()

	c := clock.FixedClocker{}
	normalUser := &entity.User{
		ID:        1,
		Name:      "taro",
		CreatedAt: c.Now(),
		UpdatedAt: c.Now(),
	}

	type mocks struct {
		pname        string
		createdAt    time.Time
		updatedAt    time.Time
		lastInsertID int64
	}

	tests := map[string]struct {
		wantUser *entity.User
		wantErr  error
		mocks    mocks
	}{
		"normal": {
			wantUser: normalUser,
			wantErr:  nil,
			mocks: mocks{
				pname:        normalUser.Name,
				createdAt:    normalUser.CreatedAt,
				updatedAt:    normalUser.UpdatedAt,
				lastInsertID: int64(normalUser.ID),
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			mockDb, mockExec, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { _ = mockDb.Close() })
			mockExec.ExpectExec(
				`INSERT INTO users \(
					name, created_at, updated_at
				\) VALUES \(\?, \?, \?\)`,
			).WithArgs(sub.mocks.pname, sub.mocks.createdAt, sub.mocks.updatedAt).
				WillReturnResult(sqlmock.NewResult(sub.mocks.lastInsertID, 1))

			xdb := sqlx.NewDb(mockDb, "mysql")
			r := &Repository{
				Clocker: c,
			}
			gotUser, gotErr := r.CreateUser(ctx, xdb, sub.mocks.pname)
			if gotErr != sub.wantErr {
				t.Fatalf("want error(including nil): %v, but got %v", sub.wantErr, gotErr)
			}
			if u := cmp.Diff(gotUser, sub.wantUser); len(u) != 0 {
				t.Errorf("differs: (-got +want)\n%s", u)
			}
		})
	}
}

func TestCreateUser_DeplicateUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// トランザクションを張ることで、このテストケースの中だけのテーブル状態にする
	tx, err := utils.OpenDBForTest(t).BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	// このテストケースが完了したらもとに戻すため、ロールバックする
	t.Cleanup(func() { _ = tx.Rollback() })

	duplicateName := randomUserName()
	fc := clock.FixedClocker{}
	_ = prepareUser(ctx, t, tx, duplicateName, fc)

	type mocks struct {
		pname string
	}

	tests := map[string]struct {
		wantUser *entity.User
		wantErr  error
		mocks    mocks
	}{
		"duplicate user error": {
			wantUser: nil,
			wantErr: apperror.REGISTER_DUPLICATE_DATA_RESTRICTED.Wrap(
				&mysql.MySQLError{Number: MYSQL_DUPLICATE_ENTRY_ERRCODE},
				fmt.Sprintf("cannot create same name user: %s", duplicateName),
			),
			mocks: mocks{
				pname: duplicateName,
			},
		},
	}

	for name, sub := range tests {
		sub := sub
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r := &Repository{Clocker: fc}
			gotUser, gotErr := r.CreateUser(ctx, tx, sub.mocks.pname)

			if e := cmp.Diff(gotErr, sub.wantErr,
				cmpopts.IgnoreFields(mysql.MySQLError{}, "SQLState"),
				cmpopts.IgnoreFields(mysql.MySQLError{}, "Message"),
			); len(e) != 0 {
				t.Errorf("differs: (-got +want)\n%s", e)
			}

			if u := cmp.Diff(gotUser, sub.wantUser); len(u) != 0 {
				t.Errorf("differs: (-got +want)\n%s", u)
			}
		})
	}
}

func prepareUser(ctx context.Context, t *testing.T, db Execer, name string, c clock.Clocker) entity.UserID {
	t.Helper()

	u := fixture.User(&entity.User{Name: name, CreatedAt: c.Now(), UpdatedAt: c.Now()})

	result, err := db.ExecContext(ctx,
		`INSERT INTO users (
			name, created_at, updated_at
		) VALUES (?, ?, ?)`, u.Name, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		t.Fatalf("insert user error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("got inserted user_id error: %v", err)
	}

	return entity.UserID(id)
}

func randomUserName() string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 15)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return fmt.Sprintf("taro_%s", string(b))
}
