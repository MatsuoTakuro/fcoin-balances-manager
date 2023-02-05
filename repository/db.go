package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MatsuoTakuro/fcoin-balances-manager/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
アプリを停止する場合に、DBとの接続も切ることができるように、db.Close()も返り値として返却する
*/
func OpenDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {

	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { _ = db.Close() }, nil
}

//go:generate go run github.com/matryer/moq -out db_mock.go . Beginner Execer Queryer

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Execer
	Queryer
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type Queryer interface {
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
}

var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.Tx)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
)
