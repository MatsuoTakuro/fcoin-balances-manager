package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/MatsuoTakuro/fcoin-balances-manager/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName,
		))

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "mysql")
}

func TruncateTables(ctx context.Context, db *sqlx.DB, tables []string) error {
	_, err := db.ExecContext(ctx, `SET FOREIGN_KEY_CHECKS=0`)
	if err != nil {
		return err
	}

	for _, v := range tables {
		_, err = db.ExecContext(ctx, fmt.Sprintf(`TRUNCATE TABLE %s`, v))
		if err != nil {
			return err
		}
	}

	_, err = db.ExecContext(ctx, `SET FOREIGN_KEY_CHECKS=1`)
	if err != nil {
		return err
	}

	return nil
}
