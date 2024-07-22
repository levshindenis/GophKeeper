package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"time"
)

// CheckUser - функция для проверки правильности логина и пароля
func (sd *Database) CheckUser(log models.Login) bool {
	var userID string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sd.DB.QueryRowContext(ctx,
		`SELECT user_id FROM users WHERE login = $1 and password = $2`,
		log.Login, log.Password).Scan(&userID); err != nil {
		return false
	}

	return true
}
