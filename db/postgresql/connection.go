package postgresql

import (
	"database/sql"
	// "os"

	_ "github.com/lib/pq" // PostgreSQL driver

	log "github.com/sirupsen/logrus"
)

// Connect to the PotgreSQL database and return the db connection
func Connect(postgresUrl string) *sql.DB {

	db, err := sql.Open("postgres", postgresUrl)
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
