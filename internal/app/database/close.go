package database

// Close - функция для закрытия БД
func (sd *Database) Close() error {
	if err := sd.DB.Close(); err != nil {
		return err
	}

	return nil
}
