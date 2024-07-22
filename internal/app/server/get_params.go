package server

import (
	"github.com/levshindenis/GophKeeper/internal/app/cloud"
	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/database"
)

func (s *Server) GetConfig() *config.Config {
	return &s.config
}

func (s *Server) GetCloud() *cloud.Cloud {
	return &s.cloud
}

func (s *Server) GetDB() *database.Database {
	return &s.db
}
