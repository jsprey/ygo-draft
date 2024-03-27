package model

import (
	_ "embed"
)

type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
}

type CardImage struct {
	ID            int    `json:"id"`
	ImageURL      string `json:"image_url"`
	ImageURLSmall string `json:"image_url_small"`
}

type Card struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Desc      string      `json:"desc"`
	Atk       int         `json:"atk"`
	Def       int         `json:"def"`
	Level     int         `json:"level"`
	Race      string      `json:"race"`
	Attribute string      `json:"attribute"`
	Sets      string      `json:"sets"`
	SetsMeta  []CardSet   `json:"card_sets"`
	Images    []CardImage `json:"card_images"`
}

// CardFilter is used to filter cards when requesting them via
type CardFilter struct {
	Types []string `json:"types"`
	Sets  []string `json:"sets"`
}

func (c *CardFilter) HasFilter() bool {
	return c.HasTypeFilter() || c.HasSetFilter()
}

func (c *CardFilter) HasTypeFilter() bool {
	if c.Types == nil {
		return false
	}

	return len(c.Types) != 0
}

func (c *CardFilter) HasSetFilter() bool {
	if c.Sets == nil {
		return false
	}

	return len(c.Sets) != 0
}

// CreateSetList creates a string list containing the identifier of all sets that are embedded in the card.
func (c *Card) CreateSetList() {
	setList := ""
	for i, set := range c.SetsMeta {
		if i != 0 {
			setList += ","
		}
		setList += set.SetName
	}
	c.Sets = setList
}

// YgoQueryGenerator is used to generate the ygo related queries send to the database client.
type YgoQueryGenerator interface {
	// SelectCardByID generate a select query for a specific card by the given id.
	SelectCardByID(id int) (string, error)
	// SelectAllCards generate a select query for all stored cards.
	SelectAllCards() (string, error)
	// SelectAllSets generate a select query for all stored sets.
	SelectAllSets() (string, error)
	// SelectSetByCode generate a select query for a given stored set.
	SelectSetByCode(cardSet string) (string, error)
	// SelectAllCardsWithFilter generate a select query for all stored cards with a given filter.
	SelectAllCardsWithFilter(filter CardFilter) (string, error)
	// InsertCard generates a query to insert a specific card into the database.
	InsertCard(card *Card) (string, error)
	// InsertSet generates a query to insert a specific card set into the database.
	InsertSet(set CardSet) (string, error)
}

// YgoClient decorates an entity to retrieve ygo data from a certain place.
type YgoClient interface {
	// GetAllCards retrieves all cards.
	GetAllCards() (*[]*Card, error)
	// GetAllSets retrieves all sets.
	GetAllSets() (*[]*CardSet, error)
	// GetAllCardsWithFilter retrieves all cards with a given filter.
	GetAllCardsWithFilter(filter CardFilter) (*[]*Card, error)
	// GetCard retrieves a api by the given id.
	GetCard(id int) (*Card, error)
	// GetSet retrieves a set by the given set code.
	GetSet(setCode string) (*CardSet, error)
}
