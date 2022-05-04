package setup

import (
	"fmt"
	"github.com/sirupsen/logrus"
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
func NewYgoDataSyncher(client *ygo.YgoClientWithCache, ygoCtx *config.YgoContext) *YgoDataSyncher {
	return &YgoDataSyncher{Client: client,
		YgoCtx: ygoCtx}
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
	logrus.Printf("Setup -> YgoDataSyncher -> Starting sync of api data...")

	cards, err := yds.Client.WebClient.GetAllCards()
	if err != nil {
		return fmt.Errorf("failed to get all cards: %w", err)
	}

	if yds.YgoCtx.Stage == config.StageDevelopment {
		logrus.Warn("[DEVELOPMENT-STAGE] Only synchronizing 20 cards")
		limitedCards := []*model.Card{}

		for i := 0; i < 20; i++ {
			card := (*cards)[i]
			limitedCards = append(limitedCards, card)
		}

		cards = &limitedCards
	}

	err = yds.Client.CacheSaver.SaveAllCards(cards)
	if err != nil {
		return fmt.Errorf("failed to save all cards: %w", err)
	}

	return nil
}
