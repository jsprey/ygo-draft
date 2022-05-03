package query

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
	"ygodraft/backend/model"
)

//go:embed templates/QuerySelectCardByID.sql
var TemplateContentSelectCardByID string

func QuerySelectCardByID(id int) (string, error) {
	idObject := struct {
		ID int `json:"id"`
	}{ID: id}

	t, err := template.New("QuerySelectCardByID").Parse(TemplateContentSelectCardByID)
	if err != nil {
		return "", fmt.Errorf("failed to parse QuerySelectCardByID template: %w", err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, &idObject)
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}

//go:embed templates/QuerySelectAllCards.sql
var TemplateContentSelectAllCards string

func QuerySelectAllCards() (string, error) {
	t, err := template.New("QuerySelectAllCards").Parse(TemplateContentSelectAllCards)
	if err != nil {
		return "", fmt.Errorf("failed to parse QuerySelectAllCards template: %w", err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, "")
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}

//go:embed templates/QueryInsertCard.sql
var TemplateContentInsertCard string

func QueryInsertCard(card *model.Card) (string, error) {
	cardCopy := *card
	cardCopy.Name = escape(cardCopy.Name)
	cardCopy.Desc = escape(cardCopy.Desc)
	cardCopy.Race = escape(cardCopy.Race)
	cardCopy.Attribute = escape(cardCopy.Attribute)
	cardCopy.Type = escape(cardCopy.Type)

	t, err := template.New("CardInsert").Parse(TemplateContentInsertCard)
	if err != nil {
		return "", fmt.Errorf("failed to parse QueryInsertCard template: %w", err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, cardCopy)
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}

func escape(input string) string {
	return fmt.Sprintf("'%s'", strings.Replace(input, "'", "''", -1))
}
