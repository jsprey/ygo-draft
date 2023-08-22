package usermgt

import (
	"fmt"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

type usermgtClient struct {
	Client         model.DatabaseClient
	QueryTemplater model.UsermgtQueryGenerator
}

func (u usermgtClient) GetUser(email string) (*model.User, error) {
	selectUserQuery, err := u.QueryTemplater.SelectUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to select select user by email query: %w", err)
	}

	userList := []*model.User{}
	err = u.Client.Select(selectUserQuery, &userList)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}

	if userList == nil || (userList != nil && len(userList) == 0) {
		return nil, model.ErrorUserDoesNotExist.WithParam(email)
	}

	return userList[0], nil
}

func NewUsermgtClient(dbClient model.DatabaseClient) (*usermgtClient, error) {
	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &usermgtClient{
		Client:         dbClient,
		QueryTemplater: queryTemplater,
	}, nil
}
