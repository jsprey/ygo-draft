package query

import (
	_ "embed"
	"fmt"
	"ygodraft/backend/model"
)

//go:embed templates/cards/QuerySelectCardByID.sql
var TemplateContentSelectCardByID string

//go:embed templates/cards/QuerySelectSetByCode.sql
var TemplateContentSelectSetByCode string

//go:embed templates/cards/QuerySelectAllCardsWithFilter.sql
var TemplateContentSelectAllCardsWithFilter string

//go:embed templates/cards/QuerySelectAllCards.sql
var TemplateContentSelectAllCards string

//go:embed templates/cards/QuerySelectAllSets.sql
var TemplateContentSelectAllSets string

//go:embed templates/cards/QueryInsertCard.sql
var TemplateContentInsertCard string

//go:embed templates/cards/QueryInsertSet.sql
var TemplateContentInsertSet string

func (sqt *sqlQueryTemplater) AddCardTemplates(templateMap *map[string]string) {
	(*templateMap)["SelectCardByID"] = TemplateContentSelectCardByID
	(*templateMap)["SelectSetByCode"] = TemplateContentSelectSetByCode
	(*templateMap)["SelectAllCardsWithFilter"] = TemplateContentSelectAllCardsWithFilter
	(*templateMap)["SelectAllCards"] = TemplateContentSelectAllCards
	(*templateMap)["SelectAllSets"] = TemplateContentSelectAllSets
	(*templateMap)["InsertCard"] = TemplateContentInsertCard
	(*templateMap)["InsertSet"] = TemplateContentInsertSet
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
