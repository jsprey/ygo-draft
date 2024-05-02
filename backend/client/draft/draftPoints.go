package draft

import (
	"fmt"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

type draftPointsClient struct {
	Client         model.DatabaseClient
	DraftClient    model.DraftClient
	QueryTemplater model.DraftPointsQueryGenerator
}

// NewDraftPointsClient creates a new instance of the draft points client.
func NewDraftPointsClient(dbClient model.DatabaseClient, draftClient model.DraftClient) (*draftPointsClient, error) {
	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &draftPointsClient{
		Client:         dbClient,
		DraftClient:    draftClient,
		QueryTemplater: queryTemplater,
	}, nil
}

func (d draftPointsClient) SetPoints(draftID int, userID int, points int) error {
	_, err := d.DraftClient.GetDraft(draftID, userID)
	if err != nil {
		return fmt.Errorf("failed to get draft: %w", err)
	}

	upsertPointsQuery, err := d.QueryTemplater.UpsertPoints(draftID, userID, points)
	if err != nil {
		return fmt.Errorf("failed to template query [UpsertPoints]: %w", err)
	}

	_, err = d.Client.Exec(upsertPointsQuery)
	if err != nil {
		return fmt.Errorf("failed to exec query [UpsertPoints]: %w", err)
	}

	return nil
}

func (d draftPointsClient) GetPoints(draftID int, userID int) (model.DraftPoints, error) {
	_, err := d.DraftClient.GetDraft(draftID, userID)
	if err != nil {
		return model.DraftPoints{}, fmt.Errorf("failed to get draft: %w", err)
	}

	selectPointsQuery, err := d.QueryTemplater.SelectPoints(draftID, userID)
	if err != nil {
		return model.DraftPoints{}, fmt.Errorf("failed to template query [SelectPoints]: %w", err)
	}

	var selectedRows []*model.DraftPoints
	err = d.Client.Select(selectPointsQuery, &selectedRows)
	if err != nil {
		return model.DraftPoints{}, fmt.Errorf("failed to select query [SelectPoints]: %w", err)
	}

	if selectedRows == nil || len(selectedRows) == 0 {
		return model.DraftPoints{}, nil
	}

	return *selectedRows[0], nil
}
