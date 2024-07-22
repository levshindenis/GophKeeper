package database

import (
	"context"
	"strings"
	"time"
)

func (sd *Database) DeleteCards(userId string, numbers []string, param string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if param != "" {
		if _, err = tx.ExecContext(ctx,
			`DELETE FROM cards WHERE user_id = $1`, userId); err != nil {
			return err
		}
	} else {
		for i := range numbers {
			arr := strings.Split(numbers[i], "///")
			if _, err = tx.ExecContext(ctx,
				`DELETE FROM cards WHERE user_id = $1 and bank = $2 and number = $3`, userId, arr[0], arr[1]); err != nil {
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
