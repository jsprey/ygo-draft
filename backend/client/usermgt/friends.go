package usermgt

import (
	"fmt"
	"ygodraft/backend/model"
)

func (u usermgtClient) GetFriends(userID int) ([]model.Friend, error) {
	query, err := u.QueryTemplater.GetFriends(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create [GetFriends] template: %w", err)
	}

	var userList []model.Friend
	err = u.Client.Select(query, &userList)
	if err != nil {
		return nil, fmt.Errorf("failed to exec get friends of user: %w", err)
	}

	if userList == nil {
		userList = []model.Friend{}
	}

	return userList, nil
}

func (u usermgtClient) GetFriendRequests(userID int) ([]model.FriendRequest, error) {
	query, err := u.QueryTemplater.GetFriendRequests(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create [GetFriendRequests] template: %w", err)
	}

	var requestList []model.FriendRequest
	err = u.Client.Select(query, &requestList)
	if err != nil {
		return nil, fmt.Errorf("failed to exec get friends of user: %w", err)
	}

	if requestList == nil {
		requestList = []model.FriendRequest{}
	}

	return requestList, nil
}

func (u usermgtClient) SetRelationshipStatus(userID int, user2ID int, status model.RelationshipStatus) error {
	if status == model.FriendStatusInvited {
		query, err := u.QueryTemplater.SetFriendRelation(userID, user2ID, model.FriendStatusInvited)
		if err != nil {
			return fmt.Errorf("failed to create [GetFriendRequests] template: %w", err)
		}

		_, err = u.Client.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to exec get friends of user: %w", err)
		}
	} else if status == model.FriendStatusFriends {
		queryFromTo, err := u.QueryTemplater.SetFriendRelation(userID, user2ID, model.FriendStatusFriends)
		if err != nil {
			return fmt.Errorf("failed to create [GetFriendRequests] template: %w", err)
		}

		_, err = u.Client.Exec(queryFromTo)
		if err != nil {
			return fmt.Errorf("failed to exec get friends of user: %w", err)
		}

		// save the relation in both ways
		queryToFrom, err := u.QueryTemplater.SetFriendRelation(user2ID, userID, model.FriendStatusFriends)
		if err != nil {
			return fmt.Errorf("failed to create [GetFriendRequests] template: %w", err)
		}

		_, err = u.Client.Exec(queryToFrom)
		if err != nil {
			return fmt.Errorf("failed to exec get friends of user: %w", err)
		}
	}

	return nil
}
