package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"time"
)

func (sd *Database) GetUserCards(login string) ([]models.Card, error) {
	var (
		items []models.Card
		item  models.Card
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := sd.DB.QueryContext(ctx,
		`Select bank, number, month, year, cvv, owner, comment, favourite from cards where user_id = $1`, login)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&item.Bank, &item.Number, &item.Month, &item.Year, &item.CVV,
			&item.Owner, &item.Comment, &item.Favourite); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
