package main

import (
	"log"

	"github.com/kacpertarka/restaurant/cmd/api"
)

func main() {

	server := api.NewHTTPServer(":8080")

	if err := server.Start(); err != nil {
		log.Fatalf("ERROR: failed to start server %v\n", err)
	}
}
