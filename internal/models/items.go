package models

import "time"

type Item struct {
	ID          string    `db:"id"`
	StartTime   time.Time `db:"start_time"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
