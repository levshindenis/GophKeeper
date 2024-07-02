package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) AddTexts(userId string, texts []models.Text) error {
	var str string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	for i := range texts {
		row := tx.QueryRowContext(ctx, `SELECT comment FROM texts WHERE user_id = $1 and name = $2`,
			userId, texts[i].Name)
		if err = row.Scan(&str); err == nil {
			continue
		}

		if _, err = tx.ExecContext(ctx,
			`INSERT INTO texts (user_id, name, description, comment, favourite) values ($1, $2, $3, $4, $5)`,
			userId, texts[i].Name, texts[i].Description, texts[i].Comment, texts[i].Favourite); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
