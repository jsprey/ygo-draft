package model

// YgoClient decorates an entity to retrieve ygo data from a certain place.
type YgoClient interface {
	// GetAllCards retrieves all cards.
	GetAllCards() (*[]*Card, error)
	// GetCard retrieves a api by the given id.
	GetCard(id int) (*Card, error)
}
