package setup

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/client/ygo"
)

// YgoDataSyncher synches the data of the web client into the local cache.
type YgoDataSyncher struct {
	Client *ygo.YgoClientWithCache
}

// NewYgoDataSyncher creates a new data syncher.
func NewYgoDataSyncher(client *ygo.YgoClientWithCache) *YgoDataSyncher {
	return &YgoDataSyncher{Client: client}
}

// Sync synchronizes all data from the web api to the local cache
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
	logrus.Debug("Setup -> YgoDataSyncher -> Starting sync of api data...")

	cards, err := yds.Client.WebClient.GetAllCards()
	if err != nil {
		return fmt.Errorf("failed to get all cards: %w", err)
	}

	err = yds.Client.CacheSaver.SaveAllCards(cards)
	if err != nil {
		return fmt.Errorf("failed to save all cards: %w", err)
	}

	return nil
}
