package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) ChangeCards(userId string, cards []models.ChCard) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range cards {
		if _, err = tx.ExecContext(ctx, `UPDATE cards SET bank = $1, number = $2, month = $3, year = $4, 
                 cvv = $5, owner = $6, comment = $7, favourite = $8 
             WHERE user_id = $9 and bank = $10 and number = $11`,
			cards[i].NewBank, cards[i].NewNumber, cards[i].NewMonth, cards[i].NewYear, cards[i].NewCVV,
			cards[i].NewOwner, cards[i].NewComment, cards[i].NewFavourite, userId, cards[i].OldBank,
			cards[i].OldNumber); err != nil {
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
