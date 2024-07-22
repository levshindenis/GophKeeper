package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"time"
)

func (sd *Database) ChangeCard(userId string, card models.ChCard) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err = tx.ExecContext(ctx, `UPDATE cards SET bank = $1, number = $2, month = $3, year = $4, cvv = $5, 
                 owner = $6, comment = $7, favourite = $8 WHERE user_id = $9 and bank = $10 and number = $11`,
		card.NewBank, card.NewNumber, card.NewMonth, card.NewYear, card.NewCVV, card.NewOwner, card.NewComment,
		card.NewFavourite, userId, card.OldBank, card.OldNumber); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, `UPDATE updates SET update_time = $1 where user_id = $2`,
		time.Now().Format(time.RFC3339), userId); err != nil {
		return err
	}

	tx.Commit()
	return nil
}
