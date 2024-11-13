package users

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpertarka/restaurant/utils"
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
	router.HandleFunc("/users", handler.createNewUser).Methods("POST")          // create new user
	router.HandleFunc("/users", handler.getAllUsers).Methods("GET")             // get all users
	router.HandleFunc("/users/{user_id}", handler.getSingleUser).Methods("GET") // get single user
}

func (handler *UserHandler) createNewUser(w http.ResponseWriter, r *http.Request) {
	/*
		Create new user account with given: email, first_name, last_name.
		Set variable such as role (default is worker), created_at (created time),
		Service should generate user_id (UUID), first login password (after first login it should be changed)

		Return map with user_id and password if user created successfully.  Otherwise, return error message.
	*/
	// get JSON payload
	var payload RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteERROR(w, err)
        return
	}
	// check if user with given email exists - worker with given email

	// if doesn't exist - create new user and return map with user_id and password
}

func (handler *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	// implement logic to get all users
}

func (handler *UserHandler) getSingleUser(w http.ResponseWriter, r *http.Request) {
	// implement logic to get single user
}
