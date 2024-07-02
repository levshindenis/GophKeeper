package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) DeleteCards(userId string, numbers []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range numbers {
		if _, err = tx.ExecContext(ctx,
			`DELETE FROM cards WHERE user_id = $1 and number = $2`, userId, numbers[i]); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
