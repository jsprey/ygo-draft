package query

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"ygodraft/backend/model"
)

func (sqt *sqlQueryTemplater) AddChallengeQueries(templateMap *map[string]string) {
	(*templateMap)["SelectOutgoingChallenges"] = templateContentSelectOutgoingChallenges
	(*templateMap)["SelectReceivedChallenges"] = templateContentSelectReceivedChallenges
	(*templateMap)["SelectChallenge"] = templateContentQuerySelectChallenge
	(*templateMap)["UpdateChallenge"] = templateContentUpdateChallenge
	(*templateMap)["InsertChallenge"] = templateContentInsertChallenge
}

//go:embed templates/challenges/QuerySelectOutgoingChallenges.sql
var templateContentSelectOutgoingChallenges string

func (sqt *sqlQueryTemplater) SelectOutgoingChallenges(challengerID int, status model.DraftChallengeStatus) (string, error) {
	templateObject := struct {
		ChallengerID int    `json:"challenger_id"`
		Status       string `json:"status"`
	}{ChallengerID: challengerID}

	if status != model.StatusAll {
		templateObject.Status = escape(string(status))
	}

	return sqt.Template("SelectOutgoingChallenges", &templateObject)
}

//go:embed templates/challenges/QuerySelectReceivedChallenges.sql
var templateContentSelectReceivedChallenges string

func (sqt *sqlQueryTemplater) SelectReceivedChallenges(receiverID int, status model.DraftChallengeStatus) (string, error) {
	templateObject := struct {
		ReceiverID int    `json:"receiver_id"`
		Status     string `json:"status"`
	}{ReceiverID: receiverID}

	if status != model.StatusAll {
		templateObject.Status = escape(string(status))
	}

	return sqt.Template("SelectReceivedChallenges", &templateObject)
}

//go:embed templates/challenges/QuerySelectChallenge.sql
var templateContentQuerySelectChallenge string

func (sqt *sqlQueryTemplater) SelectChallenge(challengeID int) (string, error) {
	templateObject := struct {
		ChallengeID int `json:"challenge_id"`
	}{ChallengeID: challengeID}

	return sqt.Template("SelectChallenge", &templateObject)
}

//go:embed templates/challenges/QueryInsertChallenge.sql
var templateContentInsertChallenge string

func (sqt *sqlQueryTemplater) InsertChallenge(challengerID int, receiverID int, settings model.DraftSettings) (string, error) {
	settingsJson, err := json.Marshal(settings)
	if err != nil {
		return "", fmt.Errorf("failed to marshal draft settings: %w", err)
	}

	templateObject := struct {
		ChallengerID int    `json:"challenger_id"`
		ReceiverID   int    `json:"receiver_id"`
		Status       string `json:"status"`
		Settings     string `json:"settings"`
	}{
		ChallengerID: challengerID,
		ReceiverID:   receiverID,
		Settings:     escape(string(settingsJson)),
		Status:       escape(string(model.StatusPending)),
	}

	return sqt.Template("InsertChallenge", &templateObject)
}

//go:embed templates/challenges/QueryUpdateChallenge.sql
var templateContentUpdateChallenge string

func (sqt *sqlQueryTemplater) UpdateChallenge(challengeID int, status model.DraftChallengeStatus) (string, error) {
	if status == model.StatusAll {
		return "", fmt.Errorf("invalid status value for update")
	}

	templateObject := struct {
		ChallengeID int    `json:"challenge_id"`
		Status      string `json:"status"`
	}{
		ChallengeID: challengeID,
		Status:      escape(string(status)),
	}

	return sqt.Template("UpdateChallenge", &templateObject)
}
