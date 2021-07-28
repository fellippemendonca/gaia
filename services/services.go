package services

import "database/sql"

// Services injector
type Services struct {
	Postgresql *sql.DB
}
