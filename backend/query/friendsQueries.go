package query

import (
	_ "embed"
	"ygodraft/backend/model"
)

func (sqt *sqlQueryTemplater) AddFriendsTemplates(templateMap *map[string]string) {
	(*templateMap)["GetFriends"] = templateContentGetFriends
	(*templateMap)["SetFriendRelation"] = templateContentSetFriendRelation
	(*templateMap)["GetFriendRequests"] = templateContentGetFriendRequests
	(*templateMap)["GetFriendRelation"] = templateContentGetFriendRelation
}

//go:embed templates/users/friends/QueryGetFriends.sql
var templateContentGetFriends string

func (sqt *sqlQueryTemplater) GetFriends(userID int) (string, error) {
	templateObject := struct {
		UserID int `json:"userID"`
	}{
		UserID: userID,
	}

	return sqt.Template("GetFriends", &templateObject)
}

//go:embed templates/users/friends/QueryGetFriendRequests.sql
var templateContentGetFriendRequests string

func (sqt *sqlQueryTemplater) GetFriendRequests(fromUseruserID int) (string, error) {
	templateObject := struct {
		UserID int `json:"userID"`
	}{UserID: fromUseruserID}

	return sqt.Template("GetFriendRequests", &templateObject)
}

//go:embed templates/users/friends/QuerySetFriendRelation.sql
var templateContentSetFriendRelation string

func (sqt *sqlQueryTemplater) SetFriendRelation(fromUserID int, toUserID int, status model.RelationshipStatus) (string, error) {
	templateObject := struct {
		FromUserID int    `json:"fromUserID"`
		ToUserID   int    `json:"toUserID"`
		Status     string `json:"status"`
	}{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Status:     escape(string(status)),
	}

	return sqt.Template("SetFriendRelation", &templateObject)
}

//go:embed templates/users/friends/QueryGetFriendRelation.sql
var templateContentGetFriendRelation string

func (sqt *sqlQueryTemplater) GetFriendRelation(fromUserID int, toUserID int) (string, error) {
	templateObject := struct {
		FromUserID int `json:"fromUserID"`
		ToUserID   int `json:"toUserID"`
	}{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
	}

	return sqt.Template("GetFriendRelation", &templateObject)
}
