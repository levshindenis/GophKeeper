package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) ChangeTexts(userId string, texts []models.ChText) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range texts {
		row := tx.QueryRowContext(ctx,
			`SELECT comment FROM texts WHERE user_id = $1 and name = $2 and description = $3 
                            and comment = $4 and favourite = $5`,
			userId, texts[i].OldName, texts[i].OldDescription, texts[i].OldComment, texts[i].OldFavourite)
		if err = row.Scan(&str); err != nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`UPDATE texts SET name = $1, description = $2, comment = $3, favourite = $4 
             	WHERE user_id = $5 and name = $6 and description = $7 and comment = $8 and favourite = $9`,
			texts[i].NewName, texts[i].NewDescription, texts[i].NewComment, texts[i].NewFavourite,
			userId, texts[i].OldName, texts[i].OldDescription, texts[i].OldComment, texts[i].OldFavourite); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
