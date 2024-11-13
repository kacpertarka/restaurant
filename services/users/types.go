package users

import "time"

/* INTERFACES */
type UserStore interface {
	CreateNewUser(CreateUserPayload) (*ReturnCreatedUserResponse, error)
}

/* STRUCTURES */
/* Payloads */
type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserPayload struct {
	Id          int64
	UserID      string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber *string // may by null
	Password    byte
	Addres      *string // may be null
	Role        string
	CreatedAt   time.Time
	IsActive    bool
}

/* Responses */
type ReturnCreatedUserResponse struct {
	UserID   string `json:"user_id"`
	Password byte   `json:"password"`
}

/* OTHERS */
type UserBase struct { // TODO: this name to change??
	ID    int64
	Email string
}
