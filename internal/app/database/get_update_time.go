package database

import (
	"context"
	"time"
)

func (sd *Database) GetUpdateTime(login string) (string, error) {
	var myTime string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx,
		`SELECT update_time FROM updates WHERE user_id = $1`,
		login)

	if err := row.Scan(&myTime); err != nil {
		return "", err
	}

	return myTime, nil
}
