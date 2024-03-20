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
