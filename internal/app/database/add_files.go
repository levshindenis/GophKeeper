package database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *Database) AddFiles(userId string, binaries []models.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO binaries (user_id, name, comment, favourite) values ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := range binaries {
		if _, err = stmt.ExecContext(ctx, userId, binaries[i].Name, binaries[i].Comment, binaries[i].Favourite); err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}
