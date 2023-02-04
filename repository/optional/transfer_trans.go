package optional

import "database/sql"

type TransferTransOpt struct {
	ID          sql.NullInt64 `db:"id"`
	FromUser    sql.NullInt64 `db:"from_user"`
	FromBalance sql.NullInt64 `db:"from_balance"`
	ToUser      sql.NullInt64 `db:"to_user"`
	ToBalance   sql.NullInt64 `db:"to_balance"`
	Amount      sql.NullInt32 `db:"amount"`
	ProcessedAt sql.NullTime  `db:"processed_at"`
}
