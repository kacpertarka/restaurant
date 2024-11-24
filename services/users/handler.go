package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpertarka/restaurant/utils"
)

type UserHandler struct {
	crud UserCRUD
}

func NewUserHandler(crud UserCRUD) *UserHandler {
	return &UserHandler{crud: crud}
}

/*
That handler is used by administration to creat new user account and manage all users -- login required, admin role required
*/
func (handler *UserHandler) RegisterRoutes(router *mux.Router) {
	// handle methods like: add new user, get all users, get single user
	router.HandleFunc("/users", handler.createNewUser).Methods("POST")              // create new user
	router.HandleFunc("/users", handler.getAllUsers).Methods("GET")                 // get all users
	router.HandleFunc("/users/{user_id}", handler.getSingleUser).Methods("GET")     // get single user
	router.HandleFunc("/users/first_login", handler.firstUserLogin).Methods("POST") // login
}

func (handler *UserHandler) firstUserLogin(w http.ResponseWriter, r *http.Request) {
	/*
		First user login. REquired old (auto generated) password and new - different than old
	*/
	// get JSON payload
	var payload FirstLoginUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteERROR(w, err)
		return
	}
	// validate new password, change it, generate JWT
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
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteERROR(w, err)
		return
	}
	// use user crud to add new user (user crud use storage)
	createUserResponse, err := handler.crud.CreateNewUser(payload)
	if err != nil {
		utils.WriteERROR(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createUserResponse)
}

func (handler *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	// implement logic to get all users
}

func (handler *UserHandler) getSingleUser(w http.ResponseWriter, r *http.Request) {
	// implement logic to get single user
}
