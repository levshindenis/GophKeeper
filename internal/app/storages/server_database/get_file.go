package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (sd *ServerDatabase) GetFile(userId string, name string) (models.File, error) {
	var (
		item models.File
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx, `SELECT comment, favourite FROM binaries WHERE user_id = $1 and name = $2`,
		userId, name)

	if err := row.Scan(&item.Comment, &item.Favourite); err != nil {
		return models.File{}, err
	}

	item.Name = name

	return item, nil
}
