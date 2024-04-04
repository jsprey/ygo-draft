package query

import (
	_ "embed"
	"strings"
	"ygodraft/backend/model"
)

func (sqt *sqlQueryTemplater) AddDraftQueries(templateMap *map[string]string) {
	(*templateMap)["InsertDraft"] = templateContentInsertDraft
	(*templateMap)["UpdateDraft"] = templateContentUpdateDraft
	(*templateMap)["SelectDraft"] = templateContentSelectDraft
	(*templateMap)["InsertDraftRound"] = templateContentInsertDraftRound
	(*templateMap)["UpdateDraftRound"] = templateContentQueryUpdateDraftRound
	(*templateMap)["InsertDraftRoundDeck"] = templateContentInsertDraftRoundDeck
}

//go:embed templates/drafts/QueryInsertDraft.sql
var templateContentInsertDraft string

func (sqt *sqlQueryTemplater) InsertDraft(challengeID int, fromUserID int, toUserID int, status model.DraftStatus) (string, error) {
	templateObject := struct {
		ChallengeID int    `json:"challenge_id"`
		FromUserID  int    `json:"from_user_id"`
		ToUserID    int    `json:"to_user_id"`
		Status      string `json:"status"`
	}{
		ChallengeID: challengeID,
		FromUserID:  fromUserID,
		ToUserID:    toUserID,
		Status:      escape(string(status)),
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

func (sqt *sqlQueryTemplater) SelectDraft(userID int) (string, error) {
	templateObject := struct {
		UserID int `json:"user_id"`
	}{
		UserID: userID,
	}

	return sqt.Template("SelectDraft", &templateObject)
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
