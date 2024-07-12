package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) DeleteCards(userId string, numbers []string) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range numbers {
		row := tx.QueryRowContext(ctx, `SELECT comment FROM cards WHERE user_id = $1 and number = $2`,
			userId, numbers[i])
		if err = row.Scan(&str); err != nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`DELETE FROM cards WHERE user_id = $1 and number = $2`, userId, numbers[i]); err != nil {
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
