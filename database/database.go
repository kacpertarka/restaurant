package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kacpertarka/restaurant/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage(config config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
