package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) ChangeCards(userId string, cards []models.ChCard) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range cards {
		row := tx.QueryRowContext(ctx,
			`SELECT comment FROM cards WHERE user_id = $1 and bank = $2 and number = $3 and month = $4 
                                    and year = $5 and cvv = $6 and owner = $7 and comment = $8 and favourite = $9`,
			userId, cards[i].OldBank, cards[i].OldNumber, cards[i].OldMonth, cards[i].OldYear,
			cards[i].OldCVV, cards[i].OldOwner, cards[i].OldComment, cards[i].OldFavourite)
		if err = row.Scan(&str); err != nil {
			continue
		}

		if _, err = tx.ExecContext(ctx, `UPDATE cards SET bank = $1, number = $2, month = $3, year = $4, 
                 cvv = $5, owner = $6, comment = $7, favourite = $8 
             WHERE user_id = $9 and bank = $10 and number = $11 and month = $12 and year = $13 and cvv = $14 
               and owner = $15 and comment = $16 and favourite = $17`,
			cards[i].NewBank, cards[i].NewNumber, cards[i].NewMonth, cards[i].NewYear, cards[i].NewCVV,
			cards[i].NewOwner, cards[i].NewComment, cards[i].NewFavourite, userId, cards[i].OldBank,
			cards[i].OldNumber, cards[i].OldMonth, cards[i].OldYear, cards[i].OldCVV, cards[i].OldOwner,
			cards[i].OldComment, cards[i].OldFavourite); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
