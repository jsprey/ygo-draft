package query

import (
	_ "embed"
	"fmt"
	"text/template"
)

//go:embed templates/users/QuerySelectUserByEmail.sql
var templateContentUsersSelectUserByEmail string

//go:embed templates/users/QueryInsertUser.sql
var templateContentUsersInsertUser string

func (sqt *sqlQueryTemplater) ParseUserTemplates() error {
	myTemplates := map[string]string{
		"SelectUserByEmail": templateContentUsersSelectUserByEmail,
		"InsertUser":        templateContentUsersInsertUser,
	}

	for templateName, templateString := range myTemplates {
		parsedTemplate, err := template.New(templateName).Funcs(customFunctions).Parse(templateString)
		if err != nil {
			return fmt.Errorf("failed to parse template [%s]: %w", templateName, err)
		}

		sqt.Templates[templateName] = parsedTemplate
	}

	return nil
}

func (sqt *sqlQueryTemplater) SelectUserByEmail(email string) (string, error) {
	idObject := struct {
		Email string `json:"email"`
	}{Email: escape(email)}

	return sqt.Template("SelectUserByEmail", &idObject)
}

func (sqt *sqlQueryTemplater) InsertUser(email string, passwordHash string, displayName string, isAdmin bool) (string, error) {
	idObject := struct {
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
		DisplayName  string `json:"display_name"`
		IsAdmin      string `json:"is_admin"`
	}{
		Email:        escape(email),
		PasswordHash: escape(passwordHash),
		DisplayName:  escape(displayName),
		IsAdmin:      escape(fmt.Sprintf("%t", isAdmin)),
	}

	return sqt.Template("InsertUser", &idObject)
}
