package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) GetLogin(userId string) (string, error) {
	var login string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx, `SELECT login FROM users WHERE user_id = $1`, userId)

	if err := row.Scan(&login); err != nil {
		return "", err
	}

	return login, nil
}
