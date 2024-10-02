package main

import (
	"log"

	"github.com/kacpertarka/restaurant/cmd/api"
	"github.com/kacpertarka/restaurant/config"
)

func main() {
	config := config.InitConfig()

	server := api.NewHTTPServer(config.PORT)

	if err := server.Start(); err != nil {
		log.Fatalf("ERROR: failed to start server %v\n", err)
	}
}
