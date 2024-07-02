package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) ListTexts(userId string) ([]string, error) {
	var (
		arr   []string
		value string
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := sd.DB.QueryContext(ctx, `SELECT name FROM texts WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			return nil, err
		}
		arr = append(arr, value)
	}

	return arr, nil
}
