package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"time"
)

func (sd *Database) AddCards(userId string, cards []models.Card) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO cards (user_id, bank, number, month, year, cvv, owner, comment, favourite)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := range cards {
		if _, err = stmt.ExecContext(ctx, userId, cards[i].Bank, cards[i].Number, cards[i].Month, cards[i].Year,
			cards[i].CVV, cards[i].Owner, cards[i].Comment, cards[i].Favourite); err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}
