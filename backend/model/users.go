package model

import (
	"time"
	"ygodraft/backend/customerrors"
)

var (
	ErrorUserDoesNotExist = customerrors.WithCode{
		Code:        "EC_User_NotFound",
		InternalMsg: "the requested user with email %s does not exist",
	}
)

// RelationshipStatus represents the state of the relationship between two users.
type RelationshipStatus string

const FriendStatusInvited RelationshipStatus = "invited"
const FriendStatusFriends RelationshipStatus = "friends"
const FriendStatusUnrelated RelationshipStatus = "unrelated"

// User represents a user in the database and is used for authentication.
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	DisplayName  string `json:"display_name"`
	IsAdmin      bool   `json:"is_admin"`
}

// Friend represents the friends of a user. Contains only relevant information about friends for a user.
type Friend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// FriendRequest represents a request of another user.
type FriendRequest struct {
	Friend
	InvitationDate time.Time `json:"invitation_date"`
}

type UsermgtClient interface {
	// GetUser returns the user with the given email. Throws an error if the user does not exist.
	GetUser(email string) (*User, error)
	// GetUser returns the user with the given id. Throws an error if the user does not exist.
	GetUserByID(id int) (*User, error)
	// CreateUser creates a new user with the given data.
	CreateUser(newUser User) error
	// DeleteUser deletes the user with the given email.
	DeleteUser(userEmail string) error
	// GetFriends returns the friends of a given user.
	GetFriends(userID int) ([]Friend, error)
	// GetFriendRequests returns the pending friend requests for the user.
	GetFriendRequests(userID int) ([]FriendRequest, error)
	// SetRelationshipStatus sets the relationship between two users.
	SetRelationshipStatus(userID int, user2ID int, status RelationshipStatus) error
	// GetRelationshipStatus retrieves the current relationship between two users.
	//GetRelationshipStatus(userID int, user2ID int) (RelationshipStatus, error)
}

type UsermgtQueryGenerator interface {
	// InsertUser creates an insert query for a new user.
	InsertUser(newUser User) (string, error)
	// SelectUserByEmail creates a select query for a specific user by the given email.
	SelectUserByEmail(email string) (string, error)
	// SelectUserByID creates a select query for a specific user by the given id.
	SelectUserByID(id int) (string, error)
	// DeleteUser creates a delete query for a specific user by the given email.
	DeleteUser(email string) (string, error)
	// GetFriends creates a select query to retrieve all friends for a certain user.
	GetFriends(userID int) (string, error)
	// GetFriendRequests creates a select query to retrieves all pending friend requests for a user.
	GetFriendRequests(userID int) (string, error)
	// SetFriendRelation creates an execution query to define the relation between two users.
	SetFriendRelation(fromUserID int, toUserID int, status RelationshipStatus) (string, error)
	// GetFriendRelation creates a select query to request the relation between two users.
	GetFriendRelation(fromUserID int, toUserID int) (string, error)
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
