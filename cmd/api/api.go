package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	addr string
}

func NewHTTPServer(addr string) *HTTPServer {
	return &HTTPServer{addr: addr}
}

func (server *HTTPServer) Start() error {
	// Implement HTTP server start logic here
	// Use gorilla/mux to create a new server instance

	router := mux.NewRouter()
	// subrouter := router.PathPrefix("/api/v1").Subrouter()
	// Register routes and handlers here

	// start server
	log.Printf("Starting server on %s", server.addr)
	return http.ListenAndServe(server.addr, router)
}
