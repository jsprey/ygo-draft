package synch

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
)

// YgoDataSyncher synches the data of the web client into the local cache.
type YgoDataSyncher struct {
	Client *ygo.YgoClientWithCache
	YgoCtx *config.YgoContext
}

// NewYgoDataSyncher creates a new data syncher.
func NewYgoDataSyncher(client *ygo.YgoClientWithCache, ygoCtx *config.YgoContext) (*YgoDataSyncher, error) {
	err := ensurePathExists(ImagesDirectoryName)
	if err != nil {
		return nil, fmt.Errorf("failed to check directory for storing images: %w", err)
	}

	return &YgoDataSyncher{Client: client,
		YgoCtx: ygoCtx}, nil
}

// Sync synchronizes all data from the web api to the local cache.
func (yds *YgoDataSyncher) Sync() error {
	logrus.Debug("Setup -> YgoDataSyncher -> Starting sync...")

	err := yds.SyncAllCards()
	if err != nil {
		return fmt.Errorf("failed to synch cards: %w", err)
	}

	return nil
}

// SyncAllCards synchronizes all api data from the web api to the local cache
func (yds *YgoDataSyncher) SyncAllCards() error {
	logrus.Printf("YgoDataSyncher -> Starting sync of api data...")

	cards, err := yds.Client.WebClient.GetAllCards()
	if err != nil {
		return fmt.Errorf("failed to get all cards: %w", err)
	}

	if yds.YgoCtx.Stage == config.StageDevelopment {
		logrus.Warn("YgoDataSyncher -> [DEVELOPMENT-STAGE] Only synchronizing 20 cards")
		limitedCards := []*model.Card{}

		for i := 0; i < 200; i++ {
			if i >= len(*cards) {
				break
			}

			card := (*cards)[i]
			limitedCards = append(limitedCards, card)
		}

		cards = &limitedCards
	}

	numberOfCards := len(*cards)
	logrus.Infof("YgoDataSyncher -> Starting the synchronization of %d cards...", numberOfCards)

	currentCard := 0
	updateTicker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-updateTicker.C:
				logrus.Infof("YgoDataSyncher -> Synched cards %d/%d", currentCard, numberOfCards)
			case <-quit:
				updateTicker.Stop()
				return
			}
		}
	}()

	for i, card := range *cards {
		_, err := yds.Client.CacheRetriever.GetCard(card.ID)
		if errors.Is(err, cache.ErrorCardDoesNotExist) {
			err = yds.Client.CacheCardSaver.SaveCard(card)
			if err != nil {
				return fmt.Errorf("failed to save card [%d]: %w", card.ID, err)
			}
		}

		//err = yds.SynchImageOfCard(card)
		//if err != nil {
		//	return fmt.Errorf("failed to synch card images: %w", err)
		//}

		err = yds.Client.CacheSetSaver.SaveSets(card.Sets)
		if err != nil {
			return fmt.Errorf("failed to synch card images: %w", err)
		}

		currentCard = i
	}

	close(quit)
	updateTicker.Stop()

	return nil
}
