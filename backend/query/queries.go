package query

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"
	"ygodraft/backend/model"
)

//go:embed templates/QuerySelectCardByID.sql
var TemplateContentSelectCardByID string

//go:embed templates/QuerySelectAllCards.sql
var TemplateContentSelectAllCards string

//go:embed templates/QueryInsertCard.sql
var TemplateContentInsertCard string

func (sqt *sqlQueryTemplater) ParseCardTemplates() error {
	myTemplates := map[string]string{
		"SelectCardByID": TemplateContentSelectCardByID,
		"SelectAllCards": TemplateContentSelectAllCards,
		"InsertCard":     TemplateContentInsertCard,
	}

	for templateName, templateString := range myTemplates {
		parsedTemplate, err := template.New(templateName).Parse(templateString)
		if err != nil {
			return fmt.Errorf("failed to parse template [%s]: %w", templateName, err)
		}

		sqt.Templates[templateName] = parsedTemplate
	}

	return nil
}

func (sqt *sqlQueryTemplater) SelectCardByID(id int) (string, error) {
	idObject := struct {
		ID int `json:"id"`
	}{ID: id}

	return sqt.Template("SelectCardByID", &idObject)
}

func (sqt *sqlQueryTemplater) SelectAllCards() (string, error) {
	return sqt.Template("SelectAllCards", "")
}

func (sqt *sqlQueryTemplater) InsertCard(card *model.Card) (string, error) {
	cardCopy := *card
	cardCopy.Name = escape(cardCopy.Name)
	cardCopy.Desc = escape(cardCopy.Desc)
	cardCopy.Race = escape(cardCopy.Race)
	cardCopy.Attribute = escape(cardCopy.Attribute)
	cardCopy.Type = escape(cardCopy.Type)

	return sqt.Template("InsertCard", &cardCopy)
}

func escape(input string) string {
	return fmt.Sprintf("'%s'", strings.Replace(input, "'", "''", -1))
}
