package domain

import "time"

// User is a registered person of this service
type User struct {
	ID            int64     `db:"id"`
	Name          string    `db:"name"`
	Email         string    `db:"email"`
	PasswordHash  string    `db:"password"`
	IsDisabled    bool      `db:"is_disabled"`
	EmailVerified bool      `db:"email_verified"`
	EmailToken    string    `db:"email_token"`
	FailedLogins  int       `db:"failed_logins"`
	LastLogin     time.Time `db:"last_login"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
