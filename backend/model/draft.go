package model

import (
	"time"
	"ygodraft/backend/customerrors"
)

var (
	ErrorDraftDoesNotExist = customerrors.WithCode{
		Code:        "EC_Draft_Does_Not_Exist",
		InternalMsg: "the requested draft with id %s does not exist",
	}
	ErrorUserAlreadyChallenged = customerrors.WithCode{
		Code:        "EC_Challenge_User_Already_Challenged",
		InternalMsg: "the receiving user already have a pending challenge from the challenger",
	}
	ErrorUserAlreadyRunningDraft = customerrors.WithCode{
		Code:        "EC_Draft_Users_Already_In_Draft",
		InternalMsg: "the users already have an unfinished draft",
	}
	ErrorDraftIsNotAChallenge = customerrors.WithCode{
		Code:        "EC_Draft_Is_Not_A_Challenge",
		InternalMsg: "the draft is not a challenge and can neither be accepted or declined",
	}
	ErrorOnlyReceivingPartyCanAcceptChallenge = customerrors.WithCode{
		Code:        "EC_Challenge_Only_Receiver_Can_Accept",
		InternalMsg: "only the receiving user can accept this challenge",
	}
	ErrorOnlyReceivingPartyCanDeclineChallenge = customerrors.WithCode{
		Code:        "EC_Challenge_Only_Receiver_Can_Decline",
		InternalMsg: "only the receiving user can decline this challenge",
	}
	ErrorUserIsNotParticipatingInDraft = customerrors.WithCode{
		Code:        "EC_User_Is_Not_Participating_In_Draft",
		InternalMsg: "you are not part of this challenge and have no access to it",
	}
)

// IsErrorCustom checks if the given error of a custom error.
func IsErrorCustom(err error, targetError customerrors.WithCode) bool {
	if err == nil {
		return false
	}

	customError, ok := err.(customerrors.WithCode)
	if !ok {
		return false
	}

	return customError.Code == targetError.Code
}

// DraftStatus determines the current state of the draft.
type DraftStatus string

const (
	// DraftStatusPending show that someone is currently challenging another player.
	DraftStatusPending DraftStatus = "pending"
	// DraftStatusDeclined show that the receiving party declined the challenge.
	DraftStatusDeclined DraftStatus = "declined"
	// DraftStatusRunning shows that someone is currently drafting another player.
	DraftStatusRunning DraftStatus = "running"
	// DraftStatusCanceled shows that someone canceled the draft preemptively.
	DraftStatusCanceled DraftStatus = "canceled"
	// DraftStatusFinished shows that the draft is finished.
	DraftStatusFinished DraftStatus = "finished"
)

// DraftRoundStatus determines the current state of the draft round.
type DraftRoundStatus string

const (
	DraftRoundStatusPreparation DraftRoundStatus = "preparation"
	DraftRoundStatusFighting    DraftRoundStatus = "fighting"
	DraftRoundStatusFinished    DraftRoundStatus = "finished"
)

// DraftMode determines the mode of the draft.
type DraftMode string

const (
	DraftModeBestOf DraftMode = "bestof"
	DraftGoalRounds DraftMode = "round"
)

// Draft contains the information for a draft.
type Draft struct {
	ID                 int           `json:"id"`
	ChallengerID       int           `json:"challenger_id"`
	ReceiverID         int           `json:"receiver_id"`
	CurrentRoundNumber int           `json:"current_round_number"`
	MaximumRoundNumber int           `json:"maximum_round_number"`
	WinnerUserID       int           `json:"winner_user_id"`
	Status             DraftStatus   `json:"status"`
	Settings           DraftSettings `json:"settings"`
	ChallengeDate      time.Time     `json:"challenge_date"`
}

// DraftSettings contains the configurable settings for a draft.
type DraftSettings struct {
	MainDeckDraws  int       `json:"main_deck_draws"`
	MainDeckSize   int       `json:"main_deck_size"`
	ExtraDeckDraws int       `json:"extra_deck_draws"`
	ExtraDeckSize  int       `json:"extra_deck_size"`
	Mode           DraftMode `json:"mode"`
	ModeValue      int       `json:"mode_value"`
	Sets           []CardSet `json:"sets"`
}

// DraftClient provides all necessary function to control and manage the drafts.
type DraftClient interface {
	// CreateDraftChallenge creates a new draft between the given users with the status model.DraftStatusPending.
	CreateDraftChallenge(challengerID int, receiverID int, settings DraftSettings) error
	// UserHaveDraftWithStatus shows if there is already a draft between two users with the given state.
	UserHaveDraftWithStatus(user1ID int, user2ID int, status DraftStatus) (bool, error)

	// AcceptDraftChallenge accepts a draft challenge and changes the status of the draft to model.DraftStatusRunning.
	AcceptDraftChallenge(draftID int, userID int) error
	// DeclineDraftChallenge declines a draft challenge and changes the status of the draft to model.DraftStatusDeclined.
	DeclineDraftChallenge(draftID int, userID int) error

	// GetDraftsWithStatus returns all drafts for the user with the given status.
	GetDraftsWithStatus(userID int, status DraftStatus) ([]Draft, error)
	// GetDraft returns the specific draft with the given id.
	GetDraft(draftID int, userID int) (Draft, error)

	// SurrenderRunningDraft surrenders the given draft and automatically makes the enemy user the winner.
	SurrenderRunningDraft(draftID int, surrenderingUserID int) error
}

// DraftQueryGenerator is responsible to generate queries related to the draft process.
type DraftQueryGenerator interface {
	// InsertDraft returns an insert query to create a new entry in the draft table.
	InsertDraft(challengerID int, receiverID int, settings DraftSettings) (string, error)

	// SelectDraft returns a select query to get a specific draft.
	SelectDraft(draftID int) (string, error)
	// SelectDraftsWithStatus returns a select query to get the drafts for the given user with a certain status.
	SelectDraftsWithStatus(userID int, status DraftStatus) (string, error)
	// SelectDraftsWithUsersAndStatus returns a select query to get all draft with a user and status filter.
	SelectDraftsWithUsersAndStatus(userID int, user2ID int, status DraftStatus) (string, error)

	// UpdateDraft returns an update query to update an entry in the draft table.
	UpdateDraft(draftID int, status DraftStatus) (string, error)
	// InsertDraftRound returns an insert query to create a new draft round.
	InsertDraftRound(draftID int, roundNumber int, status DraftRoundStatus) (string, error)
	// UpdateDraftRound returns an update query to update the draft round. When the winnerID is provided with 0, it will only update the status.
	UpdateDraftRound(draftRoundID int, winnerID int, status DraftRoundStatus) (string, error)
	// InsertDraftRoundDeck creates an insert query to create a new deck for a draft round.
	InsertDraftRoundDeck(roundID int, userID int, deck []string) (string, error)
}
