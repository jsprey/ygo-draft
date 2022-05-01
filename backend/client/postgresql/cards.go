package postgresql

import (
	"fmt"
	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/types"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"ygodraft/backend/customerrors"
	"ygodraft/backend/model"
)

const (
	TableCards = "cards"

	QueryGetAllCards = "SELECT * FROM %s"
	QueryGetCardByID = "SELECT * FROM %s WHERE id == %d"
	QuerySaveCard    = "INSERT INTO %s VALUES ?"
)

var (
	ErrorCardDoesNotExists = customerrors.WithCode{
		Code:        "mycode",
		InternalMsg: "the requested card [%s] does not exists",
	}
)

func (yc *YgoCache) createCardsTable() error {
	logrus.Debug("Cache -> Creating table for cards")

	err := yc.createTable(TableCards)
	if err != nil {
		return fmt.Errorf("failed to create table [%s] in database: %w", TableCards, err)
	}

	return nil
}

func (yc *YgoCache) GetAllCards() (*[]*model.Card, error) {
	logrus.Debug("Cache -> Retrieve all cards")

	query := fmt.Sprintf(QueryGetAllCards, TableCards)
	res, err := yc.GenjiDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query [%s]: %w", query, err)
	}
	defer res.Close()

	myCards := []*model.Card{}
	err = res.Iterate(func(d types.Document) error {
		var card model.Card
		err = document.StructScan(d, &card)
		if err != nil {
			return fmt.Errorf("failed to scan struct: %w", err)
		}

		myCards = append(myCards, &card)
		return nil
	})

	return &myCards, nil
}

func (yc *YgoCache) GetCard(id int) (*model.Card, error) {
	logrus.Debugf("Cache -> Retrieve card by id %d", id)

	query := fmt.Sprintf(QueryGetCardByID, TableCards, id)
	res, err := yc.GenjiDB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query [%s]: %w", query, err)
	}
	defer res.Close()

	var card model.Card
	err = res.Iterate(func(d types.Document) error {
		err = document.StructScan(d, &card)
		if err != nil {
			return fmt.Errorf("failed to scan struct: %w", err)
		}

		return nil
	})

	if card.ID == 0 {
		return nil, ErrorCardDoesNotExists.WithParam(strconv.Itoa(id))
	}

	return &card, nil
}

func (yc *YgoCache) SaveAllCards(cards *[]*model.Card) error {
	numberOfCards := len(*cards)
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
		if err != nil {
			return fmt.Errorf("failed to get card with id=[%d]: %w", card.ID, err)
		}

		err = yc.SaveCard(card)
		if err != nil {
			return fmt.Errorf("failed to save card [%s]: %w", card.ID, err)
		}

		currentCard = i
	}

	close(quit)

	return nil
}

func (yc *YgoCache) SaveCard(card *model.Card) error {
	logrus.Debugf("Cache -> Save card with id %d", card.ID)

	query := fmt.Sprintf(QuerySaveCard, TableCards)
	err := yc.GenjiDB.Exec(query, card)
	if err != nil {
		return fmt.Errorf("failed to exec [%s]: %w", query, err)
	}

	return nil
}
