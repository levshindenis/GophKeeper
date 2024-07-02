package server_database

import (
	"context"
	"strconv"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

// AddUser - функция для регистрации клиента
func (sd *ServerDatabase) AddUser(reg models.Register) (bool, string, error) {
	var (
		userID string
		count  int
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := sd.DB.QueryRowContext(ctx,
		`SELECT user_id FROM users WHERE login = $1`,
		reg.Login)

	if err := row.Scan(&userID); err == nil {
		return true, "", nil
	}

	row = sd.DB.QueryRowContext(ctx, `SELECT count(*) from users`)

	if err := row.Scan(&count); err != nil {
		return false, "", err
	}

	tx, err := sd.DB.Begin()
	if err != nil {
		return false, "", err
	}

	if _, err = tx.ExecContext(ctx,
		`INSERT INTO users (user_id, login, password, word) values ($1, $2, $3, $4)`,
		strconv.Itoa(count+1), reg.Login, reg.Password, reg.Word); err != nil {
		tx.Rollback()
		return false, "", err
	}

	if _, err = tx.ExecContext(ctx,
		`INSERT INTO updates (user_id, update_time) values ($1, $2)`,
		strconv.Itoa(count+1), time.Now().Format(time.RFC3339)); err != nil {
		return false, "", err
	}

	tx.Commit()
	return false, strconv.Itoa(count + 1), nil
}
