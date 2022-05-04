package cache

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"ygodraft/backend/customerrors"
	"ygodraft/backend/model"
)

var (
	ErrorCardDoesNotExist = customerrors.WithCode{
		Code:        "mycode",
		InternalMsg: "the requested api [%s] does not exists",
	}
)

func (yc *YgoCache) GetAllCards() (*[]*model.Card, error) {
	logrus.Debug("Cache -> Retrieve all cards")

	sqlQuery, err := yc.QueryTemplater.SelectAllCards()
	if err != nil {
		return nil, err
	}

	var cards []*model.Card
	err = yc.Client.Select(sqlQuery, &cards)
	if err != nil {
		return nil, fmt.Errorf("failed to scan struct: %w", err)
	}

	return &cards, nil
}

func (yc *YgoCache) GetCard(id int) (*model.Card, error) {
	logrus.Debugf("Cache -> Retrieve card by id %d", id)
	sqlQuery, err := yc.QueryTemplater.SelectCardByID(id)
	if err != nil {
		return nil, err
	}

	var cards []*model.Card
	err = yc.Client.Select(sqlQuery, &cards)
	if err != nil {
		return nil, fmt.Errorf("failed to scan struct: %w", err)
	}

	if cards == nil || (cards != nil && len(cards) == 0) {
		return nil, ErrorCardDoesNotExist.WithParam(strconv.Itoa(id))
	}

	return cards[0], nil
}

func (yc *YgoCache) SaveAllCards(cards *[]*model.Card) error {
	numberOfCards := len(*cards)
	logrus.Infof("Cache -> Starting the synchronization of %d cards...", numberOfCards)

	currentCard := 0
	updateTicker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-updateTicker.C:
				logrus.Infof("Cache -> Synched cards %d/%d", currentCard, numberOfCards)
			case <-quit:
				updateTicker.Stop()
				return
			}
		}
	}()

	for i, card := range *cards {
		_, err := yc.GetCard(card.ID)
		if errors.Is(err, ErrorCardDoesNotExist) {
			err = yc.SaveCard(card)
			if err != nil {
				return fmt.Errorf("failed to save card [%d]: %w", card.ID, err)
			}
		}

		currentCard = i
	}

	close(quit)

	return nil
}

func (yc *YgoCache) SaveCard(card *model.Card) error {
	logrus.Debugf("Cache -> Save api with id %d", card.ID)
	sqlQuery, err := yc.QueryTemplater.InsertCard(card)
	if err != nil {
		return err
	}

	_, err = yc.Client.Exec(sqlQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [%s]: %w", sqlQuery, err)
	}

	return nil
}
