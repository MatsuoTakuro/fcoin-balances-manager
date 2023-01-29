package entity

import "time"

type UserID uint64

type User struct {
	ID        UserID    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
