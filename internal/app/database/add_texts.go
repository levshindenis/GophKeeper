package database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *Database) AddTexts(userId string, texts []models.Text) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO texts (user_id, name, description, comment, favourite) values ($1,$2,$3,$4,$5)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := range texts {
		if _, err = stmt.ExecContext(ctx, userId, texts[i].Name, texts[i].Description, texts[i].Comment,
			texts[i].Favourite); err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}
