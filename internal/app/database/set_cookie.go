package database

import (
	"context"
	"time"
)

func (sd *Database) SetCookie(userId string, login string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if login != "" {
		if _, err := sd.DB.ExecContext(ctx,
			`UPDATE users set user_id = $1 where login = $2`,
			userId, login); err != nil {

			return err
		}
	} else {
		if _, err := sd.DB.ExecContext(ctx,
			`UPDATE users set user_id = $1 where user_id = $2`,
			login, userId); err != nil {

			return err
		}
	}

	return nil
}
