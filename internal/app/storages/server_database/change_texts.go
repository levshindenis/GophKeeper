package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) ChangeTexts(userId string, texts []models.ChText) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range texts {
		if _, err = tx.ExecContext(ctx,
			`UPDATE texts SET name = $1, description = $2, comment = $3, favourite = $4 
             	WHERE user_id = $5 and name = $6`,
			texts[i].NewName, texts[i].NewDescription, texts[i].NewComment, texts[i].NewFavourite,
			userId, texts[i].OldName); err != nil {
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
