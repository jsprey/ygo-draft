package model

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
