package ygo

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// Sync synchronizes all data from the web api to the local cache
func (ycwc *YgoClientWithCache) Sync() error {
	logrus.Debug("API-Client -> Starting sync...")

	err := ycwc.SyncAllCards()
	if err != nil {
		return fmt.Errorf("failed to synch cards: %w", err)
	}

	return nil
}

// SyncAllCards synchronizes all card data from the web api to the local cache
func (ycwc *YgoClientWithCache) SyncAllCards() error {
	logrus.Debug("API-Client -> Starting sync of card data...")

	cards, err := ycwc.WebClient.GetAllCards()
	if err != nil {
		return fmt.Errorf("failed to get all cards: %w", err)
	}

	err = ycwc.Cache.SaveAllCards(cards)
	if err != nil {
		return fmt.Errorf("failed to save all cards: %w", err)
	}

	return nil
}
