package users

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

/*
That handler is used by administration to creat new user account and manage all users -- login required, admin role required
*/
func (handler *UserHandler) RegisterRoutes(router *mux.Router) {
	// handle methods like: add new user, get all users, get single user
	router.HandleFunc("/users", handler.createNewUser).Methods("POST") 	// create new user
	router.HandleFunc("/users", handler.getAllUsers).Methods("GET") 	// get all users
	router.HandleFunc("/users/{user_id}", handler.getSingleUser).Methods("GET")	// get single user
}

func (handler *UserHandler) createNewUser(w http.ResponseWriter, r *http.Request) {
	// implement logic to create new user
}


func (handler *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	// implement logic to get all users
}

func (handler *UserHandler) getSingleUser(w http.ResponseWriter, r *http.Request) {
	// implement logic to get single user
}