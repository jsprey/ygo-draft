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

// CardSetSaver is responsible to save card sets.
type CardSetSaver interface {
	// SaveSets saves all given card sets into the database.
	SaveSets(cards []model.CardSet) error
}

// CardRetriever is responsible retrieve cards from any kind of source.
type CardRetriever interface {
	// GetAllCards retrieves all cards from a source.
	GetAllCards() (*[]*model.Card, error)
	// GetAllSets retrieves all sets from a source.
	GetAllSets() (*[]*model.CardSet, error)
	// GetAllCardsWithFilter retrieves all cards with a given filter.
	GetAllCardsWithFilter(filter model.CardFilter) (*[]*model.Card, error)
	// GetCard retrieves a api with a certain id from a source.
	GetCard(id int) (*model.Card, error)
}

// YgoClientWithCache is a special ygo client by retrieving the ygo data from a web and storing them persistently on
// the client.
type YgoClientWithCache struct {
	CacheCardSaver CardSaver     `json:"cache_card_saver"`
	CacheRetriever CardRetriever `json:"cache_retriever"`
	CacheSetSaver  CardSetSaver  `json:"cache_set_saver"`
	WebClient      CardRetriever `json:"web_client"`
}

func NewYgoClientWithCache(client cache.DatabaseClient) (*YgoClientWithCache, error) {
	ygoCache, err := cache.NewYgoCache(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ygo cache: %w", err)
	}

	webClient := ygoprodeck.NewYgoProDeckClient()
	return &YgoClientWithCache{
		CacheCardSaver: ygoCache,
		CacheRetriever: ygoCache,
		CacheSetSaver:  ygoCache,
		WebClient:      webClient,
	}, nil
}

func (ycwc *YgoClientWithCache) GetAllCards() (*[]*model.Card, error) {
	logrus.Debug("YGO-Client -> Retrieve all cards")
	return ycwc.CacheRetriever.GetAllCards()
}

func (ycwc *YgoClientWithCache) GetAllSets() (*[]*model.CardSet, error) {
	logrus.Debug("YGO-Client -> Retrieve all sets")
	return ycwc.CacheRetriever.GetAllSets()
}

func (ycwc *YgoClientWithCache) GetCard(id int) (*model.Card, error) {
	logrus.Debugf("YGO-Client -> Retrieve api with id [%d]", id)
	card, err := ycwc.CacheRetriever.GetCard(id)
	if errors.Is(err, cache.ErrorCardDoesNotExist) {
		webCard, webErr := ycwc.WebClient.GetCard(id)
		if webErr != nil {
			return nil, fmt.Errorf("failed to get api [%d] from web api: %w", id, webErr)
		}

		webErr = ycwc.CacheCardSaver.SaveCard(webCard)
		if webErr != nil {
			return nil, fmt.Errorf("failed to save api [%d] to cache: %w", webCard.ID, webErr)
		}

		return webCard, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get api [%d] from cache: %w", id, err)
	}

	return card, nil
}

func (ycwc *YgoClientWithCache) GetAllCardsWithFilter(filter model.CardFilter) (*[]*model.Card, error) {
	logrus.Debugf("YGO-Client -> Retrieve all cards with filter [%v]", filter)

	return ycwc.CacheRetriever.GetAllCardsWithFilter(filter)
}
