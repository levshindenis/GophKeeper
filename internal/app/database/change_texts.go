package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"time"
)

func (sd *Database) ChangeTexts(userId string, text models.ChText) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`UPDATE texts SET name = $1, description = $2, comment = $3, favourite = $4 WHERE user_id = $5 and name = $6`,
		text.NewName, text.NewDescription, text.NewComment, text.NewFavourite, userId, text.OldName); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, `UPDATE updates SET update_time = $1 where user_id = $2`,
		time.Now().Format(time.RFC3339), userId); err != nil {
		return err
	}

	tx.Commit()
	return nil
}
