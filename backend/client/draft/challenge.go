package draft

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

type challengeClient struct {
	Client         model.DatabaseClient
	QueryTemplater model.DraftChallengeQueryGenerator
}

// NewChallengeClient creates a new instance of the challenge client.
func NewChallengeClient(dbClient model.DatabaseClient) (*challengeClient, error) {
	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &challengeClient{
		Client:         dbClient,
		QueryTemplater: queryTemplater,
	}, nil
}

func (c challengeClient) GetChallenge(challengeID int) (model.DraftChallenge, error) {
	selectQuery, err := c.QueryTemplater.SelectChallenge(challengeID)
	if err != nil {
		return model.DraftChallenge{}, fmt.Errorf("failed to create [SelectChallenge] template: %w", err)
	}

	var challengeList []model.DraftChallenge
	err = c.Client.Select(selectQuery, &challengeList)
	if err != nil {
		return model.DraftChallenge{}, fmt.Errorf("failed to exec [SelectChallenge]: %w", err)
	}

	if challengeList == nil || len(challengeList) == 0 {
		return model.DraftChallenge{}, model.ErrorChallengeDoesNotExist.WithParam(string(rune(challengeID)))
	}

	return challengeList[0], nil
}

func (c challengeClient) GetChallenges(userID int, status model.DraftChallengeStatus) ([]model.DraftChallenge, error) {
	selectQuery, err := c.QueryTemplater.SelectReceivedChallenges(userID, status)
	if err != nil {
		return nil, fmt.Errorf("failed to create [SelectReceivedChallenges] template: %w", err)
	}

	var challengeList []model.DraftChallenge
	err = c.Client.Select(selectQuery, &challengeList)
	if err != nil {
		return nil, fmt.Errorf("failed to exec [SelectReceivedChallenges]: %w", err)
	}

	return challengeList, nil
}

func (c challengeClient) IsChallenging(fromUser int, toUser int) (bool, error) {
	selectQuery, err := c.QueryTemplater.SelectOutgoingChallenges(fromUser, model.StatusPending)
	if err != nil {
		return false, fmt.Errorf("failed to create [SelectOutgoingChallenges] template: %w", err)
	}

	var challengeList []model.DraftChallenge
	err = c.Client.Select(selectQuery, &challengeList)
	if err != nil {
		return false, fmt.Errorf("failed to exec [SelectOutgoingChallenges]: %w", err)
	}

	for _, challenge := range challengeList {
		if challenge.ReceiverID == toUser {
			return true, nil
		}
	}

	return false, nil
}

func (c challengeClient) ChallengeUser(fromUser int, toUser int, settings model.DraftSettings) error {
	challenging, err := c.IsChallenging(fromUser, toUser)
	if err != nil {
		return fmt.Errorf("failed to check if a user is already challenging another user: %w", err)
	}

	if challenging {
		return model.ErrorUserAlreadyChallenged.WithParam(string(rune(toUser)))
	}

	insertQuery, err := c.QueryTemplater.InsertChallenge(fromUser, toUser, settings)
	if err != nil {
		return fmt.Errorf("failed to create [InsertChallenge] template: %w", err)
	}

	_, err = c.Client.Exec(insertQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [InsertChallenge]: %w", err)
	}

	return nil
}

func (c challengeClient) AcceptChallenge(challengeID int) error {
	updateQuery, err := c.QueryTemplater.UpdateChallenge(challengeID, model.StatusAccepted)
	if err != nil {
		return fmt.Errorf("failed to create [UpdateChallenge] template: %w", err)
	}

	// todo create a draft for the players with the settings

	_, err = c.Client.Exec(updateQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateChallenge]: %w", err)
	}

	return nil
}

func (c challengeClient) DeclineChallenge(challengeID int) error {
	updateQuery, err := c.QueryTemplater.UpdateChallenge(challengeID, model.StatusDeclined)
	if err != nil {
		return fmt.Errorf("failed to create [UpdateChallenge] template: %w", err)
	}

	logrus.Printf(updateQuery)

	_, err = c.Client.Exec(updateQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateChallenge]: %w", err)
	}

	return nil
}
