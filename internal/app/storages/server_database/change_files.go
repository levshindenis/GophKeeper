package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) ChangeFiles(userId string, binaries []models.ChFile) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range binaries {
		row := tx.QueryRowContext(ctx,
			`SELECT comment FROM binaries WHERE user_id = $1 and name = $2 and comment = $3 and favourite = $4`,
			userId, binaries[i].OldName, binaries[i].OldComment, binaries[i].OldFavourite)
		if err = row.Scan(&str); err != nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`UPDATE binaries SET name = $1, comment = $2, favourite = $3 
                WHERE user_id = $4 and name = $5 and comment = $6 and favourite = $7`,
			binaries[i].NewName, binaries[i].NewComment, binaries[i].NewFavourite, userId, binaries[i].OldName,
			binaries[i].OldComment, binaries[i].OldFavourite); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
