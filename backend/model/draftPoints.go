package model

// DraftPoints contains the current points for a user to a certain draft.
type DraftPoints struct {
	ID      int `json:"id"`
	DraftID int `json:"draft_id"`
	UserID  int `json:"user_id"`
	Points  int `json:"points"`
}

// DraftPointsClient provides the functionality to manage the points in a draft.
type DraftPointsClient interface {
	// SetPoints sets the points for a certain user.
	SetPoints(draftID int, userID int, points int) error
	// GetPoints retrieves the points for a certain user.
	GetPoints(draftID int, userID int) (DraftPoints, error)
}

// DraftPointsQueryGenerator is responsible to generate queries related to the draft points.
type DraftPointsQueryGenerator interface {
	// UpsertPoints creates a upsert query to create/modify the points of a user.
	UpsertPoints(draftID int, userID int, points int) (string, error)
	// SelectPoints creates a select query to retrieve the points for a user.
	SelectPoints(draftID int, userID int) (string, error)
}
