package database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *Database) GetUserTexts(login string) ([]models.Text, error) {
	var (
		items []models.Text
		item  models.Text
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := sd.DB.QueryContext(ctx,
		`Select name, description, comment, favourite from texts where user_id = $1`, login)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&item.Name, &item.Description, &item.Comment, &item.Favourite); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
