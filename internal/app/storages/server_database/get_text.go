package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) GetText(userId string, name string) (models.Text, error) {
	var (
		item models.Text
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx, `SELECT description, comment, favourite FROM texts WHERE user_id = $1 and name = $2`,
		userId, name)

	if err := row.Scan(&item.Description, &item.Comment, &item.Favourite); err != nil {
		return models.Text{}, err
	}

	item.Name = name

	return item, nil
}
