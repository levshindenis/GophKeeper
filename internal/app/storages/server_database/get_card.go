package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) GetCard(userId string, number string) (models.Card, error) {
	var (
		item models.Card
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx, `SELECT bank, month, year, cvv, owner, comment, favourite FROM cards 
                            		WHERE user_id = $1 and number = $2`, userId, number)

	if err := row.Scan(&item.Bank, &item.Month, &item.Year, &item.CVV, &item.Owner, &item.Comment, &item.Favourite); err != nil {
		return models.Card{}, err
	}

	item.Number = number

	return item, nil
}
