package server

import (
	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/storages/cloud"
	"github.com/levshindenis/GophKeeper/internal/app/storages/cookie"
	"github.com/levshindenis/GophKeeper/internal/app/storages/server_database"
)

func (s *Server) GetConfig() *config.Config {
	return &s.config
}

func (s *Server) GetCookie() *cookie.Cookie {
	return &s.cookie
}

func (s *Server) GetCloud() *cloud.Cloud {
	return &s.cloud
}

func (s *Server) GetDB() *server_database.ServerDatabase {
	return &s.db
}
