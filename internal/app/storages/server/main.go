package server

import (
	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/storages/cloud"
	"github.com/levshindenis/GophKeeper/internal/app/storages/cookie"
	"github.com/levshindenis/GophKeeper/internal/app/storages/server_database"
)

type Server struct {
	config config.Config
	cookie cookie.Cookie
	cloud  cloud.Cloud
	db     server_database.ServerDatabase
}
