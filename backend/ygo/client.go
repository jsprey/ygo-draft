package ygo

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/cache"
	"ygodraft/backend/model"
	"ygodraft/backend/ygoprodeck"
)

// YgoClientWithCache is a special ygo client by retrieving the ygo data from a web and storing them persistently on
// the client.
type YgoClientWithCache struct {
	Cache     model.YgoClient `json:"cache,omitempty"`
	WebClient model.YgoClient `json:"web_client,omitempty"`
}

func NewYgoClientWithCache() (*YgoClientWithCache, error) {
	ygoCache, err := cache.NewYgoCache("mydata.db")
	if err != nil {
		return nil, fmt.Errorf("failed to create new ygo cache: %w", err)
	}

	webClient := ygoprodeck.NewYgoProDeckClient()
	return &YgoClientWithCache{
		Cache:     ygoCache,
		WebClient: webClient,
	}, nil
}

// Close closes the connection to the internal database.
func (ycwc *YgoClientWithCache) Close() error {
	return ycwc.Cache.Close()
}

func (ycwc *YgoClientWithCache) GetAllCards() (*[]*model.Card, error) {
	logrus.Debug("YGO-Client -> Retrieve all cards")
	return ycwc.Cache.GetAllCards()
}

func (ycwc *YgoClientWithCache) GetCard(id int) (*model.Card, error) {
	logrus.Debugf("YGO-Client -> Retrieve card with id [%d]", id)
	card, err := ycwc.Cache.GetCard(id)
	if errors.Is(err, cache.ErrorCardDoesNotExists) {
		webCard, webErr := ycwc.WebClient.GetCard(id)
		if webErr != nil {
			return nil, fmt.Errorf("failed to get card [%d] from web api: %w", id, webErr)
		}

		webErr = ycwc.Cache.SaveCard(webCard)
		if webErr != nil {
			return nil, fmt.Errorf("failed to save card [%d] to cache: %w", webCard.ID, webErr)
		}

		return webCard, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get card [%d] from cache: %w", id, err)
	}

	return card, nil
}

func (ycwc *YgoClientWithCache) SaveAllCards(_ *[]*model.Card) error {
	return fmt.Errorf("invalid operation")
}

func (ycwc YgoClientWithCache) SaveCard(_ *model.Card) error {
	return fmt.Errorf("operation not supported")
}
