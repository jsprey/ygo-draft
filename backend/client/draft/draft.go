package draft

import (
	"fmt"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

type draftClient struct {
	Client         model.DatabaseClient
	QueryTemplater model.UsermgtQueryGenerator
}

func NewDraftClient(dbClient model.DatabaseClient) (*draftClient, error) {
	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &draftClient{
		Client:         dbClient,
		QueryTemplater: queryTemplater,
	}, nil
}
