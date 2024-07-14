package server_database

import (
	"context"
	"time"
)

func (sd *ServerDatabase) ListFavourites(userId string) ([]string, error) {
	var (
		arr    []string
		value  string
		value1 string
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	arr = append(arr, "Texts:")

	rows, err := sd.DB.QueryContext(ctx, `SELECT name FROM texts WHERE user_id = $1 and favourite = $2`,
		userId, true)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			return nil, err
		}
		arr = append(arr, "    "+value)
	}

	arr = append(arr, "Files:")

	rows, err = sd.DB.QueryContext(ctx, `SELECT name FROM binaries WHERE user_id = $1 and favourite = $2`,
		userId, true)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			return nil, err
		}
		arr = append(arr, "    "+value)
	}

	arr = append(arr, "Cards:")

	rows, err = sd.DB.QueryContext(ctx, `SELECT bank, number FROM cards WHERE user_id = $1 and favourite = $2`,
		userId, true)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&value, &value1); err != nil {
			return nil, err
		}
		arr = append(arr, "    "+value+" "+value1)
	}

	return arr, nil
}
