package database

import (
	"context"
	"time"
)

func (sd *Database) GetLogin(userId string) (string, error) {
	var login string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sd.DB.QueryRowContext(ctx, `SELECT login FROM users where user_id = $1`, userId).Scan(&login); err != nil {
		return "", err
	}

	return login, nil
}
