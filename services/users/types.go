package users

import "time"

/* INTERFACES */
type UserStore interface {
	// CreateNewUser(CreateUserPayload) (*ReturnCreatedUserResponse, error)
	CreateNewUser(CreateUserPayload) (*ReturnCreatedUserResponse, error)
	GetUserByEmail(string) (*UserBase, error)
	IsUserExists(string) bool
}

/* STRUCTURES */
/* Payloads */
type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserPayload struct {
	// Id          int64
	UserID    string
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Role      uint8
	CreatedAt time.Time
	IsActive  bool
}

/* Responses */
type ReturnCreatedUserResponse struct {
	// return new user_id and raw password to could next login with user_id and password
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

/* OTHERS */
type UserBase struct { // TODO: this name to change??
	ID    int64
	Email string
}
