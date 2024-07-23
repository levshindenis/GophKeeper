package database

import (
	"context"
	"strings"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *Database) GetCard(userId string, number string) (models.Card, error) {
	var (
		item models.Card
	)

	bank := strings.Split(number, "///")[0]
	num := strings.Split(number, "///")[1]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx, `SELECT bank, number, month, year, cvv, owner, comment, favourite FROM cards 
                            		WHERE user_id = $1 and bank = $2 and number = $3`, userId, bank, num)

	if err := row.Scan(&item.Bank, &item.Number, &item.Month, &item.Year, &item.CVV, &item.Owner, &item.Comment,
		&item.Favourite); err != nil {
		return models.Card{}, err
	}

	return item, nil
}
