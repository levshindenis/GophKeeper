package database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *Database) ChangeFile(userId string, binary models.ChFile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err = tx.ExecContext(ctx,
		`UPDATE binaries SET name = $1, comment = $2, favourite = $3  WHERE user_id = $4 and name = $5`,
		binary.NewName, binary.NewComment, binary.NewFavourite, userId, binary.OldName); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, `UPDATE updates SET update_time = $1 where user_id = $2`,
		time.Now().Format(time.RFC3339), userId); err != nil {
		return err
	}

	tx.Commit()
	return nil
}
