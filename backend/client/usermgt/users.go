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

func (u usermgtClient) CreateUser(newUser model.User) error {
	insertUserQurery, err := u.QueryTemplater.InsertUser(newUser)
	if err != nil {
		return fmt.Errorf("failed to select select user by email query: %w", err)
	}

	_, err = u.Client.Exec(insertUserQurery)
	if err != nil {
		return fmt.Errorf("failed to exec insert new user: %w", err)
	}

	return nil
}

func (u usermgtClient) DeleteUser(userEmail string) error {
	deleteUserQuery, err := u.QueryTemplater.DeleteUser(userEmail)
	if err != nil {
		return fmt.Errorf("failed to select select user by email query: %w", err)
	}

	_, err = u.Client.Exec(deleteUserQuery)
	if err != nil {
		return fmt.Errorf("failed to exec delete user: %w", err)
	}

	return nil
}

func (u usermgtClient) GetUser(email string) (*model.User, error) {
	selectUserQuery, err := u.QueryTemplater.SelectUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to select select user by email query: %w", err)
	}

	var userList []*model.User
	err = u.Client.Select(selectUserQuery, &userList)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}

	if userList == nil || (userList != nil && len(userList) == 0) {
		return nil, model.ErrorUserDoesNotExist.WithParam(email)
	}

	return userList[0], nil
}

func (u usermgtClient) GetUserByID(userID int) (*model.User, error) {
	selectUserQuery, err := u.QueryTemplater.SelectUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to select select user by email query: %w", err)
	}

	var userList []*model.User
	err = u.Client.Select(selectUserQuery, &userList)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}

	if userList == nil || (userList != nil && len(userList) == 0) {
		return nil, model.ErrorUserDoesNotExist.WithParam(string(rune(userID)))
	}

	return userList[0], nil
}
