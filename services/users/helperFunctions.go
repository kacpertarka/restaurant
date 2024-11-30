package users

import (
	"database/sql"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const passwordChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func hashPassword(rawPassword string) ([]byte, error) {
	/*Hash given password and return password in bytes*/
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 12)
	if err != nil {
		return []byte{}, err
	}
	return hashedPassword, nil
}

func isPasswordValid(rawPassword string, dbPassword []byte) bool {
	/*Validate given password*/
	err := bcrypt.CompareHashAndPassword(dbPassword, []byte(rawPassword))
	return err == nil
}

func generateFirstLoginPassword() string {
	/*Generate simple 24 chars password */
	rawPasswordByte := make([]byte, 24)
	var randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range rawPasswordByte {
		rawPasswordByte[i] = passwordChars[randomSource.Intn(len(passwordChars))]
	}
	return string(rawPasswordByte)
}

func scanRowIntoUserBase(rows *sql.Rows) (*UserBase, error) {
	user := new(UserBase)
	if err := rows.Scan(
		&user.ID,
		&user.UserID,
		&user.Email,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func scanRowIntoUser(rows *sql.Rows) (User, error) {
	user := new(User)
	if err := rows.Scan(
		&user.UserID,
		&user.Email,
        &user.FirstName,
        &user.LastName,
	); err != nil {
		return User{}, err
	}
	return *user, nil
}

func comparedPasswords(oldPassword, newPassword string) bool {
	// compare two of passwords - used only when user first login
	return oldPassword == newPassword
}
