package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) AddCards(userId string, cards []models.Card) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range cards {
		row := tx.QueryRowContext(ctx, `SELECT comment FROM cards WHERE user_id = $1 and bank = $2 and number = $3`,
			userId, cards[i].Bank, cards[i].Number)
		if err = row.Scan(&str); err == nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`INSERT INTO cards 
    			(user_id, bank, number, month, year, cvv, owner, comment, favourite) 
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			userId, cards[i].Bank, cards[i].Number, cards[i].Month, cards[i].Year, cards[i].CVV, cards[i].Owner,
			cards[i].Comment, cards[i].Favourite); err != nil {
			tx.Rollback()
			return err
		}

		if _, err = tx.ExecContext(ctx, `UPDATE updates SET update_time = $1 where user_id = $2`,
			time.Now().Format(time.RFC3339), userId); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
