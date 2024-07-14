package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) ChangeFiles(userId string, binaries []models.ChFile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range binaries {
		if _, err = tx.ExecContext(ctx,
			`UPDATE binaries SET name = $1, comment = $2, favourite = $3 
                WHERE user_id = $4 and name = $5 and comment = $6 and favourite = $7`,
			binaries[i].NewName, binaries[i].NewComment, binaries[i].NewFavourite, userId, binaries[i].OldName); err != nil {
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
