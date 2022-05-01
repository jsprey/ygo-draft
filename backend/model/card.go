package model

import (
	"fmt"
	"strings"
)

type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
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
	Sets      []CardSet   `json:"card_sets"`
	Images    []CardImage `json:"card_images"`
}

func (c *Card) QueryInsert() string {
	name := strings.Replace(c.Name, "'", "''", -1)
	desc := strings.Replace(c.Desc, "'", "''", -1)
	return fmt.Sprintf("INSERT INTO public.cards (\"id\",\"name\",\"type\",\"desc\",\"atk\",\"def\",\"level\",\"race\",\"attribute\") VALUES (%d,'%s','%s','%s',%d,%d,%d,'%s','%s')",
		c.ID, name, c.Type, desc, c.Atk, c.Def, c.Level, c.Race, c.Attribute)
}

func (c *Card) QuerySelect() string {
	return fmt.Sprintf("SELECT (\"id\",\"name\") FROM public.cards WHERE \"id\"=%d", c.ID)
}
