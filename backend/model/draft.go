package model

import (
	"time"
	"ygodraft/backend/customerrors"
)

var (
	ErrorChallengeDoesNotExist = customerrors.WithCode{
		Code:        "EC_Challenge_Not_Exist",
		InternalMsg: "the requested challenge with id %s does not exist",
	}
	ErrorUserAlreadyChallenged = customerrors.WithCode{
		Code:        "EC_Challenge_User_Already_Challenged",
		InternalMsg: "the user %s already has a pending challenge",
	}
)

// IsErrorChallengeDoesNotExist checks if the given error is of type ErrorUserDoesNotExist.
func IsErrorChallengeDoesNotExist(err error) bool {
	if err == nil {
		return false
	}

	customError, ok := err.(customerrors.WithCode)
	if !ok {
		return false
	}

	return customError.Code == ErrorChallengeDoesNotExist.Code
}

// IsErrorUserAlreadyChallenged checks if the given error is of type ErrorUserAlreadyChallenged.
func IsErrorUserAlreadyChallenged(err error) bool {
	if err == nil {
		return false
	}

	customError, ok := err.(customerrors.WithCode)
	if !ok {
		return false
	}

	return customError.Code == ErrorUserAlreadyChallenged.Code
}

// DraftChallengeStatus determines the current state of the challenge
type DraftChallengeStatus string

const (
	StatusAll      DraftChallengeStatus = "all"
	StatusPending  DraftChallengeStatus = "pending"
	StatusAccepted DraftChallengeStatus = "accepted"
	StatusDeclined DraftChallengeStatus = "declined"
)

// DraftMode determines the mode of the draft.
type DraftMode string

const (
	DraftModeBestOf DraftMode = "bestof"
	DraftGoalRounds DraftMode = "round"
)

// DraftSettings contains the configurable settings for a draft.
type DraftSettings struct {
	MainDeckDraws  int       `json:"main_deck_draws"`
	ExtraDeckDraws int       `json:"extra_deck_draws"`
	Mode           DraftMode `json:"mode"`
	ModeValue      int       `json:"modeValue"`
	Sets           []CardSet `json:"sets"`
}

// DraftChallenge contains the necessary information about a challenge to a draft.
type DraftChallenge struct {
	ID            int                  `json:"id"`
	ChallengerID  int                  `json:"challenger_id"`
	ReceiverID    int                  `json:"receiver_id"`
	ChallengeDate time.Time            `json:"challenge_date"`
	Status        DraftChallengeStatus `json:"status"`
	Settings      DraftSettings        `json:"settings"`
}

// DraftChallengeClient provides all necessary functions to control and manage the draft challenges.
type DraftChallengeClient interface {
	// GetChallenge Returns a specific challenges with the given id.
	GetChallenge(challengeID int) (DraftChallenge, error)
	// GetChallenges Returns all challenges from the given user.
	GetChallenges(userID int, status DraftChallengeStatus) ([]DraftChallenge, error)
	// IsChallenging returns true when a user has already a 'pending' challenge for another user.
	IsChallenging(fromUser int, toUser int) (bool, error)
	// ChallengeUser initiates a challenge between two users.
	ChallengeUser(fromUser int, toUser int, settings DraftSettings) error
	// AcceptChallenge accepts the given challenge for the receiving user.
	AcceptChallenge(challengeID int) error
	// DeclineChallenge declines the given challenge for the receiving user.
	DeclineChallenge(challengeID int) error
}

// DraftChallengeQueryGenerator is responsible to generate queries related to the draft challenges process.
type DraftChallengeQueryGenerator interface {
	// SelectChallenge returns a select query to select a specific challenge.
	SelectChallenge(challengeID int) (string, error)
	// SelectOutgoingChallenges returns a select query to select the outgoing challenges of a specific user.
	SelectOutgoingChallenges(challengerID int, status DraftChallengeStatus) (string, error)
	// SelectReceivedChallenges returns a select query to select the received challenges of a specific user.
	SelectReceivedChallenges(receiverID int, status DraftChallengeStatus) (string, error)
	// InsertChallenge returns an insert query to create a new challenge.
	InsertChallenge(challengerID int, receiverID int, settings DraftSettings) (string, error)
	// UpdateChallenge returns an update query to update challenges.
	UpdateChallenge(challengeID int, status DraftChallengeStatus) (string, error)
}

// DraftClient provides all necessary functions to control and manage the drafts between users.
type DraftClient interface {
}

// DraftQueryGenerator is responsible to generate queries related to the draft and drafting process.
type DraftQueryGenerator interface {
}
