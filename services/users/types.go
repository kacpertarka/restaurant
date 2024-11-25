package users

import "time"

/* INTERFACES */
type UserStore interface { // TODO: is this interface necessary???
	// CreateNewUser(CreateUserPayload) (*ReturnCreatedUserResponse, error)
	CreateNewUser(CreateUserPayload) (*ReturnCreatedUserResponse, error)
	GetUserByEmail(email string) (*UserBase, error)
	IsUserExists(email string) bool
	FirstChangePassword(userID int64, newPassword []byte) error
}

/* STRUCTURES */
/* Payloads */
type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type FirstLoginUserPayload struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
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

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

/* OTHERS */
type UserBase struct { // TODO: this name to change??
	ID     int64
	UserID string
	Email  string
}
