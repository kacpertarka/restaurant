package users

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) GetUserByEmail(email string) (*UserBase, error) {
	// query to get user by given email and return only UserBase(id, email)
	rows, err := store.db.Query("SELECT id, email FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(UserBase)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

/* helper function */
func scanRowIntoUser(rows *sql.Rows) (*UserBase, error) {
	user := new(UserBase)
	if err := rows.Scan(
		user.ID,
		user.Email,
	); err != nil {
		return nil, err
	}
	return user, nil
}
