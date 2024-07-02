package server_database

import (
	"database/sql"
)

type ServerDatabase struct {
	DB *sql.DB
}
