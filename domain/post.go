package domain

import "time"

// Post is a text message the User can publish
type Post struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
