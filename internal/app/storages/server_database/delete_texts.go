package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) DeleteTexts(userId string, names []string) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range names {
		row := tx.QueryRowContext(ctx, `SELECT comment FROM texts WHERE user_id = $1 and name = $2`,
			userId, names[i])
		if err = row.Scan(&str); err != nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`DELETE FROM texts WHERE user_id = $1 and name = $2`, userId, names[i]); err != nil {
			tx.Rollback()
			return err
		}

		if _, err = tx.ExecContext(ctx, `UPDATE updates SET update_time = $1 where user_id = $2`,
			time.Now().Format(time.RFC3339), userId); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
