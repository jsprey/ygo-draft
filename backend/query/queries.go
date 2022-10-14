package query

import (
	_ "embed"
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"ygodraft/backend/model"
)

//go:embed templates/QuerySelectCardByID.sql
var TemplateContentSelectCardByID string

//go:embed templates/QuerySelectSetByCode.sql
var TemplateContentSelectSetByCode string

//go:embed templates/QuerySelectAllCardsWithFilter.sql
var TemplateContentSelectAllCardsWithFilter string

//go:embed templates/QuerySelectAllCards.sql
var TemplateContentSelectAllCards string

//go:embed templates/QuerySelectAllSets.sql
var TemplateContentSelectAllSets string

//go:embed templates/QueryInsertCard.sql
var TemplateContentInsertCard string

//go:embed templates/QueryInsertSet.sql
var TemplateContentInsertSet string

var fns = template.FuncMap{
	"notLast": func(x int, a interface{}) bool {
		return x < reflect.ValueOf(a).Len()-1
	},
}

func (sqt *sqlQueryTemplater) ParseCardTemplates() error {
	myTemplates := map[string]string{
		"SelectCardByID":           TemplateContentSelectCardByID,
		"SelectAllCards":           TemplateContentSelectAllCards,
		"SelectSetByCode":          TemplateContentSelectSetByCode,
		"SelectAllSets":            TemplateContentSelectAllSets,
		"SelectAllCardsWithFilter": TemplateContentSelectAllCardsWithFilter,
		"InsertCard":               TemplateContentInsertCard,
		"InsertSet":                TemplateContentInsertSet,
	}

	for templateName, templateString := range myTemplates {
		parsedTemplate, err := template.New(templateName).Funcs(fns).Parse(templateString)
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

func (sqt *sqlQueryTemplater) SelectSetByCode(code string) (string, error) {
	templateObject := struct {
		SetCode string `json:"code"`
	}{SetCode: code}

	return sqt.Template("SelectSetByCode", &templateObject)
}

func (sqt *sqlQueryTemplater) SelectAllCardsWithFilter(filter model.CardFilter) (string, error) {
	for i, filterType := range filter.Types {
		filter.Types[i] = escape(filterType)
	}
	for i, filterSet := range filter.Sets {
		preparedFilterSet := fmt.Sprintf("%%%s%%", filterSet)
		filter.Sets[i] = escape(preparedFilterSet)
	}

	return sqt.Template("SelectAllCardsWithFilter", &filter)
}

func (sqt *sqlQueryTemplater) SelectAllCards() (string, error) {
	return sqt.Template("SelectAllCards", "")
}

func (sqt *sqlQueryTemplater) SelectAllSets() (string, error) {
	return sqt.Template("SelectAllSets", "")
}

func (sqt *sqlQueryTemplater) InsertCard(card *model.Card) (string, error) {
	cardCopy := *card
	cardCopy.Name = escape(cardCopy.Name)
	cardCopy.Desc = escape(cardCopy.Desc)
	cardCopy.Race = escape(cardCopy.Race)
	cardCopy.Attribute = escape(cardCopy.Attribute)
	cardCopy.Type = escape(cardCopy.Type)
	cardCopy.Sets = escape(cardCopy.Sets)

	return sqt.Template("InsertCard", &cardCopy)
}

func (sqt *sqlQueryTemplater) InsertSet(set model.CardSet) (string, error) {
	setCopy := set
	setCopy.SetCode = escape(setCopy.SetCode)
	setCopy.SetName = escape(setCopy.SetName)
	setCopy.SetRarity = escape(setCopy.SetRarity)
	setCopy.SetRarityCode = escape(setCopy.SetRarityCode)

	return sqt.Template("InsertSet", &setCopy)
}

func escape(input string) string {
	return fmt.Sprintf("'%s'", strings.Replace(input, "'", "''", -1))
}
