package model

import (
	"ygodraft/backend/customerrors"
)

var (
	ErrorUserDoesNotExist = customerrors.WithCode{
		Code:        "EC_User_NotFound",
		InternalMsg: "the requested user with email %s does not exist",
	}
)

// User represents a user in the database and is used for authentication.
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	DisplayName  string `json:"display_name"`
	IsAdmin      bool   `json:"is_admin"`
}

type UsermgtClient interface {
	// GetUser returns the user with the given email. Throws an error if the user does not exist.
	GetUser(email string) (*User, error)
	// CreateUser creates a new user with the given data.
	CreateUser(newUser User) error
	// DeleteUser deletes the user with the given email.
	DeleteUser(userEmail string) error
}

type UsermgtQueryGenerator interface {
	// InsertUser creates an insert query for a new user.
	InsertUser(newUser User) (string, error)
	// SelectUserByEmail creates a select query for a specific user by the given email.
	SelectUserByEmail(email string) (string, error)
	// DeleteUser creates a delete query for a specific user by the given email.
	DeleteUser(email string) (string, error)
}

// IsErrorUserDoesNotExist checks if the given error is of type ErrorUserDoesNotExist.
func IsErrorUserDoesNotExist(err error) bool {
	if err == nil {
		return false
	}

	customError, ok := err.(customerrors.WithCode)
	if !ok {
		return false
	}

	return customError.Code == ErrorUserDoesNotExist.Code
}
