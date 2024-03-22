package query

import (
	_ "embed"
	"fmt"
	"ygodraft/backend/model"
)

func (sqt *sqlQueryTemplater) AddUserTemplates(templateMap *map[string]string) {
	(*templateMap)["SelectUserByID"] = templateContentUsersSelectUserByID
	(*templateMap)["SelectUserByEmail"] = templateContentUsersSelectUserByEmail
	(*templateMap)["InsertUser"] = templateContentUsersInsertUser
	(*templateMap)["DeleteUser"] = templateContentUsersDeleteUser
}

//go:embed templates/users/QuerySelectUserByID.sql
var templateContentUsersSelectUserByID string

func (sqt *sqlQueryTemplater) SelectUserByID(id int) (string, error) {
	idObject := struct {
		ID int `json:"id"`
	}{ID: id}

	return sqt.Template("SelectUserByID", &idObject)
}

//go:embed templates/users/QuerySelectUserByEmail.sql
var templateContentUsersSelectUserByEmail string

func (sqt *sqlQueryTemplater) SelectUserByEmail(email string) (string, error) {
	idObject := struct {
		Email string `json:"email"`
	}{Email: escape(email)}

	return sqt.Template("SelectUserByEmail", &idObject)
}

//go:embed templates/users/QueryInsertUser.sql
var templateContentUsersInsertUser string

func (sqt *sqlQueryTemplater) InsertUser(newUser model.User) (string, error) {
	idObject := struct {
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
		DisplayName  string `json:"display_name"`
		IsAdmin      string `json:"is_admin"`
	}{
		Email:        escape(newUser.Email),
		PasswordHash: escape(newUser.PasswordHash),
		DisplayName:  escape(newUser.DisplayName),
		IsAdmin:      escape(fmt.Sprintf("%t", newUser.IsAdmin)),
	}

	return sqt.Template("InsertUser", &idObject)
}

//go:embed templates/users/QueryDeleteUser.sql
var templateContentUsersDeleteUser string

func (sqt *sqlQueryTemplater) DeleteUser(email string) (string, error) {
	idObject := struct {
		Email string `json:"email"`
	}{
		Email: escape(email),
	}

	return sqt.Template("DeleteUser", &idObject)
}
