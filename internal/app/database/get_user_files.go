package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"time"
)

func (sd *Database) GetUserFiles(login string) ([]models.File, error) {
	var (
		items []models.File
		item  models.File
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := sd.DB.QueryContext(ctx,
		`Select name, comment, favourite from binaries where user_id = $1`, login)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&item.Name, &item.Comment, &item.Favourite); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
