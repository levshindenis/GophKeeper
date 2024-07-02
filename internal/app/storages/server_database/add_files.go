package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) AddFiles(userId string, binaries []models.File) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range binaries {
		row := tx.QueryRowContext(ctx, `SELECT comment FROM binaries WHERE user_id = $1 and name = $2`,
			userId, binaries[i].Name)
		if err = row.Scan(&str); err == nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`INSERT INTO binaries (user_id, name, comment, favourite) values ($1, $2, $3, $4)`,
			userId, binaries[i].Name, binaries[i].Comment, binaries[i].Favourite); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
