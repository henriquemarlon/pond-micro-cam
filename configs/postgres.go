package configs

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SetupPostgres() *sqlx.DB {
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		log.Printf("POSTGRES_URL is not set")
	}

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Printf("%v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
	}

	return db
}
