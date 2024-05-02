package query

import (
	_ "embed"
)

func (sqt *sqlQueryTemplater) AddDraftPointsQueries(templateMap *map[string]string) {
	(*templateMap)["SelectPoints"] = templateContentSelectPoints
	(*templateMap)["UpsertPoints"] = templateContentUpsertPoints
}

//go:embed templates/drafts/points/QueryUpsertPoints.sql
var templateContentUpsertPoints string

func (sqt *sqlQueryTemplater) UpsertPoints(draftID int, userID int, points int) (string, error) {
	templateObject := struct {
		DraftID int `json:"draft_id"`
		UserID  int `json:"user_id"`
		Points  int `json:"points"`
	}{
		DraftID: draftID,
		UserID:  userID,
		Points:  points,
	}

	return sqt.Template("UpsertPoints", &templateObject)
}

//go:embed templates/drafts/points/QuerySelectPoints.sql
var templateContentSelectPoints string

func (sqt *sqlQueryTemplater) SelectPoints(draftID int, userID int) (string, error) {
	templateObject := struct {
		DraftID int `json:"draft_id"`
		UserID  int `json:"user_id"`
	}{
		DraftID: draftID,
		UserID:  userID,
	}

	return sqt.Template("SelectPoints", &templateObject)
}
