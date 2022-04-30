package model

// YgoClient decorates an entity to retrieve ygo data from a certain place.
type YgoClient interface {
	// GetAllCards retrieves all cards.
	GetAllCards() (*[]*Card, error)
	// GetCard retrieves a card by the given id.
	GetCard(id int) (*Card, error)
	// SaveAllCards stores all cards locally.
	SaveAllCards(cards *[]*Card) error
	// Close closes any constructed connections.
	Close() error
}
