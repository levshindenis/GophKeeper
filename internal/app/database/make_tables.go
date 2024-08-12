package database

import (
	"context"
	"time"
)

// MakeTables - функция для создания таблиц при запуске сервера
func (sd *Database) MakeTables(param string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if param != "" {
		if _, err = tx.ExecContext(ctx,
			`CREATE TABLE IF NOT EXISTS users(user_id text, login text, password text, word text)`); err != nil {
			return err
		}
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS texts(user_id text, name text, description text, comment text, favourite text)`); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS binaries(user_id text, name text, comment text, favourite text)`); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS cards(user_id text, bank text, number text, 
			month text, year text, cvv text, owner text, comment text, favourite text)`); err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS updates(user_id text, update_time text)`); err != nil {
		return err
	}

	tx.Commit()
	return nil
}
