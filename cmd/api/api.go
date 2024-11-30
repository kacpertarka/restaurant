package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpertarka/restaurant/services/orders"
	"github.com/kacpertarka/restaurant/services/users"
)

type HTTPServer struct {
	addr string
	db   *sql.DB
}

func NewHTTPServer(addr string, db *sql.DB) *HTTPServer {
	return &HTTPServer{
		addr: addr,
		db:   db,
	}
}

func (server *HTTPServer) Start() error {
	// Implement HTTP server start logic here
	// Use gorilla/mux to create a new server instance

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// Register routes and handlers here
	orderHandler := orders.NewOrderHandler(server.db)
	orderHandler.RegisterRoutes(subrouter)

	// Register user storage and handlees
	userStore := users.NewStore(server.db)
	userCRUD := users.NewUserCRUD(*userStore)
	userHandler := users.NewUserHandler(userCRUD)
	userHandler.RegisterRoutes(subrouter)

	// start server
	log.Printf("Starting server on %s", server.addr)
	return http.ListenAndServe(server.addr, router)
}
