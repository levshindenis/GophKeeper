package server_database

import (
	"context"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

// CheckUser - функция для проверки правильности логина и пароля
func (sd *ServerDatabase) CheckUser(log models.Login) (string, error) {
	var userID string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx,
		`SELECT user_id FROM users WHERE login = $1 and password = $2`,
		log.Login, log.Password)

	if err := row.Scan(&userID); err != nil {
		return "", err
	}

	return userID, nil
}
