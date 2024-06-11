package models

import "database/sql"

type Alert struct {
	ID        string       `db:"id"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt sql.NullTime `db:"created_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`

	UserId string `db:"user_id"`
	Alert  string `db:"alert"`
}

type Alerts struct {
	Alerts []string
	UserId string
}
