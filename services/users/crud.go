package users

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// user roles
type Roles uint8

const (
	MANAGER Roles = iota // usefull for admin page maybe?
	ADMIN                // can everything
	WORKER               // TODO: think!!
)

type UserCRUD struct {
	store UserStore
}

func NewUserCRUD(store UserStore) UserCRUD {
	return UserCRUD{store: store}
}

func (crud *UserCRUD) CreateNewUser(userPayload RegisterUserPayload) (*ReturnCreatedUserResponse, error) {
	/*
		Create new user - generate first login passwrod, generate user_id (uuid)
	*/
	// check if user with given email exists - worker with given email
	userEmail := userPayload.Email
	if !crud.store.IsUserExists(userEmail) {
		return nil, fmt.Errorf("user with given email address: %v already exists", userEmail)
	}

	// generate uder_id
	userID := uuid.New().String()

	// generate and hash first password
	generatedRawPassword := generateFirstLoginPassword()
	hashedPassword, err := hashPassword(generatedRawPassword)
	if err != nil {
		return nil, err
	}

	// create user create payload and use userStorage to add data to database
	createUserPayload := CreateUserPayload{
		UserID:    userID,
		FirstName: userPayload.FirstName,
		LastName:  userPayload.LastName,
		Email:     userEmail,
		Password:  hashedPassword,
		Role:      uint8(WORKER),
		CreatedAt: time.Now(),
		IsActive:  false,
	}

	createUserResponse, err := crud.store.CreateNewUser(createUserPayload)
	if err != nil {
		return nil, err
	}

	// we want to return raw password to new user can chang it then
	createUserResponse.Password = generatedRawPassword

	return createUserResponse, nil
}
