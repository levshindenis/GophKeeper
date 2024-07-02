package server_database

// Close - функция для закрытия БД
func (sd *ServerDatabase) Close() error {
	if err := sd.DB.Close(); err != nil {
		return err
	}

	return nil
}
