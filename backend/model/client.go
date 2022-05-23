package model

// YgoClient decorates an entity to retrieve ygo data from a certain place.
type YgoClient interface {
	// GetAllCards retrieves all cards.
	GetAllCards() (*[]*Card, error)
	// GetAllCardsWithFilter retrieves all cards with a given filter.
	GetAllCardsWithFilter(filter CardFilter) (*[]*Card, error)
	// GetCard retrieves a api by the given id.
	GetCard(id int) (*Card, error)
}

// CardFilter is used to filter cards when requesting them via
type CardFilter struct {
	Types []string `json:"types"`
	Sets  []string `json:"sets"`
}
