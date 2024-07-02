package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) DeleteFiles(userId string, names []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range names {
		if _, err = tx.ExecContext(ctx,
			`DELETE FROM binaries WHERE user_id = $1 and name = $2`, userId, names[i]); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
