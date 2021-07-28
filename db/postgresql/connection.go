package postgresql

import (
	_ "github.com/lib/pq" // PostgreSQL driver

	"database/sql"

	log "github.com/sirupsen/logrus"
)

// Connect to the PotgreSQL database and return the db connection
func Connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://administrator:password@192.168.178.5:5432/payment_service?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("PostgreSQL Database Connected Successfully!")
	return db
}
