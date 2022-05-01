package cache

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"ygodraft/backend/customerrors"
	"ygodraft/backend/model"
)

const (
	TableCards = "cards"

	QueryGetAllCards = "SELECT * FROM public.cards"
	QueryGetCardByID = "SELECT * FROM public.cards WHERE id == %d"
)

var (
	ErrorCardDoesNotExist = customerrors.WithCode{
		Code:        "mycode",
		InternalMsg: "the requested api [%s] does not exists",
	}
)

func (yc *YgoCache) GetAllCards() (*[]*model.Card, error) {
	logrus.Debug("Cache -> Retrieve all cards")

	//query := fmt.Sprintf(QueryGetAllCards, TableCards)
	//res, err := yc.GenjiDB.Query(query)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to query [%s]: %w", query, err)
	//}
	//defer res.Close()
	//
	//myCards := []*model.Card{}
	//err = res.Iterate(func(d types.Document) error {
	//	var card model.Card
	//	err = document.StructScan(d, &card)
	//	if err != nil {
	//		return fmt.Errorf("failed to scan struct: %w", err)
	//	}
	//
	//	myCards = append(myCards, &card)
	//	return nil
	//})
	//
	//return &myCards, nil
	return nil, nil
}

func (yc *YgoCache) GetCard(id int) (*model.Card, error) {
	logrus.Debugf("Cache -> Retrieve card by id %d", id)
	stubCard := model.Card{ID: id}

	query := fmt.Sprintf(stubCard.QuerySelect())

	var test []*struct {
		ID   int
		Name string
	}
	//var card []*model.Card
	err := yc.Client.Select(&test, query)

	//TODO [jsprey] du wolltest hier weitermachen. Get Card implementieren
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrorCardDoesNotExist.WithParam(strconv.Itoa(id))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to scan struct: %w", err)
	}

	return nil, nil
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
		if errors.Is(err, ErrorCardDoesNotExist) {
			err = yc.SaveCard(card)
			if err != nil {
				return fmt.Errorf("failed to save card [%s]: %w", card.ID, err)
			}
		}

		currentCard = i
	}

	close(quit)

	return nil
}

func (yc *YgoCache) SaveCard(card *model.Card) error {
	logrus.Debugf("Cache -> Save api with id %d", card.ID)

	query := fmt.Sprintf(card.QueryInsert())
	err := yc.Client.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to exec [%s]: %w", query, err)
	}

	return nil
}
