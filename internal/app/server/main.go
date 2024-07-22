package server

import (
	"github.com/levshindenis/GophKeeper/internal/app/cloud"
	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/database"
)

type Server struct {
	config config.Config
	cloud  cloud.Cloud
	db     database.Database
}
