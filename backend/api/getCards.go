package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ygodraft/backend/model"
)

const GetCardQueryParamID = "id"
const GetRandomCardsQueryParamSize = "size"
const GetRandomCardsQueryParamSets = "sets"
const GetRandomCardsQueryParamTypes = "types"

type CardRetrieveHandler struct {
	YGOClient model.YgoClient
}

type getCardResponse struct {
	CardIds []int `json:"card_ids"`
	Number  int   `json:"number"`
}

func (crh *CardRetrieveHandler) GetCards(ctx *gin.Context) {
	cards, err := crh.YGOClient.GetAllCards()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := getCardResponse{
		Number:  len(*cards),
		CardIds: make([]int, len(*cards)),
	}
	for i, card := range *cards {
		response.CardIds[i] = card.ID
	}

	ctx.JSONP(http.StatusOK, response)
}

func (crh *CardRetrieveHandler) GetCard(ctx *gin.Context) {
	queryID, err := strconv.Atoi(ctx.Param(GetCardQueryParamID))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
	}

	logrus.Debugf("API-Handler -> Retrieve api [%d]", queryID)

	card, err := crh.YGOClient.GetCard(queryID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	logrus.Debug(card)

	ctx.JSONP(http.StatusOK, card)
}

func (crh *CardRetrieveHandler) GetRandomCard(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve random card")

	cards, err := crh.YGOClient.GetAllCards()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cardsBox := *cards

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randomCard := cardsBox[rand.Intn(len(cardsBox))]

	logrus.Debug(randomCard)

	ctx.JSONP(http.StatusOK, randomCard)
}

func (crh *CardRetrieveHandler) GetRandomCards(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve random deck")

	numberOfCards, cardFilter, err := getRandomCardsCheckQueryAttributes(ctx)
	if err != nil {
		return
	}

	cards, err := crh.YGOClient.GetAllCardsWithFilter(cardFilter)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if *cards != nil {
		cardsBox := *cards

		randomDeck := make([]*model.Card, numberOfCards)
		for i := 0; i < numberOfCards; i++ {
			randomDeck[i] = cardsBox[rand.Intn(len(cardsBox))]
		}
		randomCardsResponse := struct {
			Cards []*model.Card `json:"cards"`
		}{Cards: randomDeck}

		ctx.JSONP(http.StatusOK, randomCardsResponse)
	} else {
		emptyCards := []*model.Card{}
		emptyResponse := struct {
			Cards []*model.Card `json:"cards"`
		}{Cards: emptyCards}

		ctx.JSONP(http.StatusOK, emptyResponse)
	}
}

func getRandomCardsCheckQueryAttributes(ctx *gin.Context) (int, model.CardFilter, error) {
	filter := model.CardFilter{}
	numberOfCards := 0

	numberOfCardsRaw, ok := ctx.Request.URL.Query()[GetRandomCardsQueryParamSize]
	if ok {
		numberOfCardsParsed, err := strconv.Atoi(numberOfCardsRaw[0])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("invalid decksize, only integers are supported"))
			return 0, model.CardFilter{}, fmt.Errorf("bad request")
		}

		numberOfCards = numberOfCardsParsed
	}

	typeListRaw, ok := ctx.Request.URL.Query()[GetRandomCardsQueryParamTypes]
	if ok {
		typeList := strings.Split(typeListRaw[0], ",")
		filter.Types = typeList
	}

	setsListRaw, ok := ctx.Request.URL.Query()[GetRandomCardsQueryParamSets]
	if ok {
		setsList := strings.Split(setsListRaw[0], ",")
		filter.Sets = setsList
	}

	return numberOfCards, filter, nil
}
