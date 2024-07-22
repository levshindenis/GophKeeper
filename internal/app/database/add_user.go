package database

import (
	"context"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

// AddUser - функция для регистрации клиента
func (sd *Database) AddUser(reg models.Register) (string, bool, error) {
	var (
		helpStr string
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sd.DB.QueryRowContext(ctx,
		`SELECT password FROM users WHERE login = $1`,
		reg.Login).Scan(&helpStr); err == nil {
		return "", true, nil
	}

	helpStr, err := tools.GenerateCookie(strconv.Itoa(rand.Intn(100)))
	if err != nil {
		return "", false, err
	}

	if _, err = sd.DB.ExecContext(ctx,
		`INSERT INTO users (user_id, login, password, word) values ($1, $2, $3, $4)`,
		helpStr, reg.Login, reg.Password, reg.Word); err != nil {
		return "", false, err
	}

	return helpStr, false, nil
}
