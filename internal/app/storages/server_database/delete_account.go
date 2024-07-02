package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) DeleteAccount(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, `DELETE FROM users WHERE user_id = $1`, userId); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, `DELETE FROM texts WHERE user_id = $1`, userId); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, `DELETE FROM binaries WHERE user_id = $1`, userId); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, `DELETE FROM cards WHERE user_id = $1`, userId); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, `DELETE FROM updates WHERE user_id = $1`, userId); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
