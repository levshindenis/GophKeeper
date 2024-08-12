package database

import (
	"context"
	"time"
)

func (sd *Database) DeleteTexts(userId string, names []string, param string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if param != "" {
		if _, err = tx.ExecContext(ctx,
			`DELETE FROM texts WHERE user_id = $1`, userId); err != nil {
			return err
		}
	} else {
		for i := range names {
			if _, err = tx.ExecContext(ctx,
				`DELETE FROM texts WHERE user_id = $1 and name = $2`, userId, names[i]); err != nil {
				return err
			}

			if _, err = tx.ExecContext(ctx, `UPDATE updates SET update_time = $1 where user_id = $2`,
				time.Now().Format(time.RFC3339), userId); err != nil {
				return err
			}
		}
	}

	tx.Commit()
	return nil
}
