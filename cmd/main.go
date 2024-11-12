package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kacpertarka/restaurant/cmd/api"
	"github.com/kacpertarka/restaurant/config"
	"github.com/kacpertarka/restaurant/database"
)

func main() {
	// load config variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := config.Envs

	// postgres database connection
	db := database.NewPostgresStorage(config)
	database.Ping(db)
	defer db.Close()
	log.Println("DB connected")

	// run API server
	server := api.NewHTTPServer(config.PORT, db)
	if err := server.Start(); err != nil {
		log.Fatalf("ERROR: failed to start server %v\n", err)
	}
}
