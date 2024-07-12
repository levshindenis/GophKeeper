package server

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/storages/cookie"
	"github.com/levshindenis/GophKeeper/internal/app/storages/server_database"
)

// Init -  функция для инициализации параметров сервера
func (s *Server) Init(conf config.Config) error {
	s.config = conf
	s.cookie = cookie.Cookie{Data: make(map[string]string)}

	//if err := s.cloud.Init(); err != nil {
	//	return err
	//}

	newDB, err := sql.Open("pgx", conf.GetDBAddress())
	if err != nil {
		return err
	}

	s.db = server_database.ServerDatabase{DB: newDB}
	if err = s.db.MakeTables(); err != nil {
		return err
	}

	return nil
}
