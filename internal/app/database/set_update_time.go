package database

import (
	"context"
	"time"
)

func (sd *Database) SetUpdateTime(userId string, updTime string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := sd.DB.ExecContext(ctx,
		`UPDATE updates set update_time = $1 where user_id = $2`,
		updTime, userId); err != nil {

		return err
	}
	return nil
}
