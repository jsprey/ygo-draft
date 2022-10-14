package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/model"
)

func SetupAPI(router *gin.RouterGroup, client *ygo.YgoClientWithCache) error {
	cardRetriever := CardRetrieveHandler{
		YGOClient: client,
	}

	router.GET("cards", cardRetriever.GetCards)
	router.GET("cards/:id", cardRetriever.GetCard)
	router.GET("randomCard", cardRetriever.GetRandomCard)
	router.GET("randomCards", cardRetriever.GetRandomCards)
	router.GET("sets", cardRetriever.GetSets)
	router.GET("sets/:code", cardRetriever.GetSet)
	router.GET("sets/:code/cards", cardRetriever.GetSetCards)

	return nil
}

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
	queryID, err := strconv.Atoi(ctx.Param("id"))
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
