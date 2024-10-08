package database

import (
	"context"
	"time"
)

func (sd *Database) ListCards(userId string) ([]string, error) {
	var (
		arr    []string
		bank   string
		number string
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := sd.DB.QueryContext(ctx, `SELECT bank, number FROM cards WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&bank, &number); err != nil {
			return nil, err
		}
		arr = append(arr, bank+"///"+number)
	}

	return arr, nil
}
