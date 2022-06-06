package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"ygodraft/backend/model"
)

const GetRandomCardsQueryParamSize = "size"
const GetRandomCardsQueryParamSets = "sets"
const GetRandomCardsQueryParamTypes = "types"

func (crh *CardRetrieveHandler) GetRandomCards(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve random deck")

	numberOfCards, cardFilter, err := getRandomCardsCheckQueryAttributes(ctx)
	if err != nil {
		return
	}

	cards, err := crh.YGOClient.GetAllCardsWithFilter(cardFilter)
	if err != nil {
		_ = ctx.AbortWithError(500, err)
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

		ctx.JSONP(200, randomCardsResponse)
	} else {
		emptyCards := []*model.Card{}
		emptyResponse := struct {
			Cards []*model.Card `json:"cards"`
		}{Cards: emptyCards}

		ctx.JSONP(200, emptyResponse)
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
