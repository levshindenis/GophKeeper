package database

import (
	"context"
	"time"
)

func (sd *Database) AddUpdateTime(userId string, updTime string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := sd.DB.ExecContext(ctx,
		`INSERT INTO updates (user_id, update_time) values ($1, $2)`,
		userId, updTime); err != nil {

		return err
	}
	return nil
}
