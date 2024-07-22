package database

import (
	"context"
	"time"
)

func (sd *Database) CheckCookie(userId string) bool {
	var helpStr string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sd.DB.QueryRowContext(ctx,
		`SELECT login FROM users WHERE user_id = $1`,
		userId).Scan(&helpStr); err != nil {
		return false
	}

	return true
}
