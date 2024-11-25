package users

import (
	"database/sql"
	"errors"
)

var notExistingUser = errors.New("user not found")

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) CreateNewUser(userPayload CreateUserPayload) (*ReturnCreatedUserResponse, error) {
	/*Add new user data do database*/
	_, err := store.db.Exec(
		"INSERT INTO users (user_id, first_name, last_name, email, password, role, created_at, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING user_id",
		userPayload.UserID, userPayload.FirstName, userPayload.LastName, userPayload.Email, userPayload.Password, userPayload.Role, userPayload.CreatedAt, userPayload.IsActive)
	if err != nil {
		return nil, err
	}
	return &ReturnCreatedUserResponse{UserID: userPayload.UserID}, nil
}

func (store *Store) GetUserByEmail(email string) (*UserBase, error) {
	// query to get user by given email and return only UserBase(id, email)
	rows, err := store.db.Query("SELECT id, user_id, email FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	user := new(UserBase)
	for rows.Next() {
		user, err = scanRowIntoUserBase(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, notExistingUser
	}
	return user, nil
}

func (store *Store) IsUserExists(email string) bool {
	/*Check if user exist by email*/
	// TODO: do not ignore other errors here
	_, err := store.GetUserByEmail(email)
	return err == notExistingUser
}

func (store *Store) FirstChangePassword(userID int64, newPassword []byte) error {
	// change password where id = user.id + activate account
	_, err := store.db.Exec("UPDATE users SET password = $1, is_active = true WHERE id = $2", newPassword, userID)
	if err != nil {
		return err
	}
	// change of passwrd went successfully
	return nil
}
