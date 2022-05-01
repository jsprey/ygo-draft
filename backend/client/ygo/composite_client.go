package ygo

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/client/ygoprodeck"
	"ygodraft/backend/model"
)

// CardSaver is responsible to save cards for later use.
type CardSaver interface {
	// SaveAllCards saves all given cards.
	SaveAllCards(cards *[]*model.Card) error
	// SaveCard saves a given api.
	SaveCard(card *model.Card) error
}

// CardRetriever is responsible retrieve cards from any kind of source.
type CardRetriever interface {
	// GetAllCards retrieves all cards from a source.
	GetAllCards() (*[]*model.Card, error)
	// GetCard retrieves a api with a certain id from a source.
	GetCard(id int) (*model.Card, error)
}

// YgoClientWithCache is a special ygo client by retrieving the ygo data from a web and storing them persistently on
// the client.
type YgoClientWithCache struct {
	CacheSaver     CardSaver     `json:"cache_saver"`
	CacheRetriever CardRetriever `json:"cache_retriever"`
	WebClient      CardRetriever `json:"web_client"`
}

func NewYgoClientWithCache(client cache.PostgresClient) (*YgoClientWithCache, error) {
	ygoCache, err := cache.NewYgoCache(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ygo cache: %w", err)
	}

	webClient := ygoprodeck.NewYgoProDeckClient()
	return &YgoClientWithCache{
		CacheSaver:     ygoCache,
		CacheRetriever: ygoCache,
		WebClient:      webClient,
	}, nil
}

func (ycwc *YgoClientWithCache) GetAllCards() (*[]*model.Card, error) {
	logrus.Debug("YGO-Client -> Retrieve all cards")
	return ycwc.CacheRetriever.GetAllCards()
}

func (ycwc *YgoClientWithCache) GetCard(id int) (*model.Card, error) {
	logrus.Debugf("YGO-Client -> Retrieve api with id [%d]", id)
	card, err := ycwc.CacheRetriever.GetCard(id)
	if errors.Is(err, cache.ErrorCardDoesNotExist) {
		webCard, webErr := ycwc.WebClient.GetCard(id)
		if webErr != nil {
			return nil, fmt.Errorf("failed to get api [%d] from web api: %w", id, webErr)
		}

		webErr = ycwc.CacheSaver.SaveCard(webCard)
		if webErr != nil {
			return nil, fmt.Errorf("failed to save api [%d] to cache: %w", webCard.ID, webErr)
		}

		return webCard, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get api [%d] from cache: %w", id, err)
	}

	return card, nil
}
