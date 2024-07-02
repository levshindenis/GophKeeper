package server_database

import (
	"context"
	"time"
)

// MakeTables - функция для создания таблиц при запуске сервера
func (sd *ServerDatabase) MakeTables() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := sd.DB.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS users(user_id text, login text, password text, word text)`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS texts(user_id text, name text, description text, comment text, favourite boolean)`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS binaries(user_id text, name text, comment text, favourite boolean)`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS cards(user_id text, bank text, number text, 
			month int, year int, cvv int, owner text, comment text, favourite boolean)`); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS updates(user_id text, update_time timestamp with time zone)`); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
