package query

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"ygodraft/backend/model"
)

func (sqt *sqlQueryTemplater) AddDraftQueries(templateMap *map[string]string) {
	(*templateMap)["InsertDraft"] = templateContentInsertDraft
	(*templateMap)["UpdateDraft"] = templateContentUpdateDraft
	(*templateMap)["SelectDraft"] = templateContentSelectDraft
	(*templateMap)["SelectDraftsWithStatus"] = templateContentSelectDraftWithStatus
	(*templateMap)["SelectDraftsWithUsersAndStatus"] = templateContentSelectDraftsWithUsersAndStatus
	(*templateMap)["InsertDraftRound"] = templateContentInsertDraftRound
	(*templateMap)["UpdateDraftRound"] = templateContentQueryUpdateDraftRound
	(*templateMap)["InsertDraftRoundDeck"] = templateContentInsertDraftRoundDeck
}

//go:embed templates/drafts/QueryInsertDraft.sql
var templateContentInsertDraft string

func (sqt *sqlQueryTemplater) InsertDraft(challengerID int, receiverID int, settings model.DraftSettings) (string, error) {
	settingsJson, err := json.Marshal(settings)
	if err != nil {
		return "", fmt.Errorf("failed to marshal settings: %w", err)
	}

	templateObject := struct {
		ChallengerID       int    `json:"challenger_id"`
		ReceiverID         int    `json:"receiver_id"`
		MaximumRoundNumber int    `json:"maximum_round_number"`
		Status             string `json:"status"`
		Settings           string `json:"settings"`
	}{
		ChallengerID:       challengerID,
		ReceiverID:         receiverID,
		MaximumRoundNumber: settings.ModeValue,
		Status:             escape(string(model.DraftStatusPending)),
		Settings:           escape(string(settingsJson)),
	}

	return sqt.Template("InsertDraft", &templateObject)
}

//go:embed templates/drafts/QueryUpdateDraft.sql
var templateContentUpdateDraft string

func (sqt *sqlQueryTemplater) UpdateDraft(draftID int, status model.DraftStatus) (string, error) {
	templateObject := struct {
		DraftID int    `json:"draft_id"`
		Status  string `json:"status"`
	}{
		DraftID: draftID,
		Status:  escape(string(status)),
	}

	return sqt.Template("UpdateDraft", &templateObject)
}

//go:embed templates/drafts/QuerySelectDraft.sql
var templateContentSelectDraft string

func (sqt *sqlQueryTemplater) SelectDraft(draftID int) (string, error) {
	templateObject := struct {
		DraftID int `json:"user_id"`
	}{
		DraftID: draftID,
	}

	return sqt.Template("SelectDraft", &templateObject)
}

//go:embed templates/drafts/QuerySelectDraftsWithStatus.sql
var templateContentSelectDraftWithStatus string

func (sqt *sqlQueryTemplater) SelectDraftsWithStatus(userID int, status model.DraftStatus) (string, error) {
	templateObject := struct {
		UserID int    `json:"user_id"`
		Status string `json:"status"`
	}{
		UserID: userID,
		Status: escape(string(status)),
	}

	return sqt.Template("SelectDraftsWithStatus", &templateObject)
}

//go:embed templates/drafts/QuerySelectDraftsWithUsersAndStatus.sql
var templateContentSelectDraftsWithUsersAndStatus string

func (sqt *sqlQueryTemplater) SelectDraftsWithUsersAndStatus(userID int, user2ID int, status model.DraftStatus) (string, error) {
	templateObject := struct {
		UserID  int    `json:"user_id"`
		User2ID int    `json:"user_2_id"`
		Status  string `json:"status"`
	}{
		UserID:  userID,
		User2ID: user2ID,
		Status:  escape(string(status)),
	}

	return sqt.Template("SelectDraftsWithUsersAndStatus", &templateObject)
}

//go:embed templates/drafts/rounds/QueryInsertDraftRound.sql
var templateContentInsertDraftRound string

func (sqt *sqlQueryTemplater) InsertDraftRound(draftID int, roundNumber int, status model.DraftRoundStatus) (string, error) {
	templateObject := struct {
		DraftID     int    `json:"draft_id"`
		RoundNumber int    `json:"round_number"`
		Status      string `json:"status"`
	}{
		DraftID:     draftID,
		RoundNumber: roundNumber,
		Status:      escape(string(status)),
	}

	return sqt.Template("InsertDraftRound", &templateObject)
}

//go:embed templates/drafts/rounds/QueryUpdateDraftRound.sql
var templateContentQueryUpdateDraftRound string

func (sqt *sqlQueryTemplater) UpdateDraftRound(draftRoundID int, winnerID int, status model.DraftRoundStatus) (string, error) {
	templateObject := struct {
		RoundID      int    `json:"round_id"`
		WinnerUserID int    `json:"winner_user_id"`
		Status       string `json:"status"`
	}{
		RoundID:      draftRoundID,
		WinnerUserID: winnerID,
		Status:       escape(string(status)),
	}

	return sqt.Template("UpdateDraftRound", &templateObject)
}

//go:embed templates/drafts/deck/QueryInsertDraftRoundDeck.sql
var templateContentInsertDraftRoundDeck string

func (sqt *sqlQueryTemplater) InsertDraftRoundDeck(roundID int, userID int, deck []string) (string, error) {
	templateObject := struct {
		RoundID int    `json:"round_id"`
		UserID  int    `json:"user_id"`
		Deck    string `json:"deck"`
	}{
		RoundID: roundID,
		UserID:  userID,
		Deck:    escape(strings.Join(deck, ",")),
	}

	return sqt.Template("InsertDraftRoundDeck", &templateObject)
}
