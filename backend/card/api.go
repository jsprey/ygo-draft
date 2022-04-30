package card

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
)

func SetupAPI(router *gin.RouterGroup, ctx *config.YGOContext) error {
	cardRetriever := CardRetrieveHandler{
		YGOClient: ctx.DataClient,
	}

	router.GET("cards", cardRetriever.GetCards)
	router.GET("cards/:id", cardRetriever.GetCard)
	router.GET("randomCard", cardRetriever.GetRandomCard)

	return nil
}

type CardRetrieveHandler struct {
	YGOClient model.YgoClient
}

func (crh *CardRetrieveHandler) GetCards(ctx *gin.Context) {
	cards, err := crh.YGOClient.GetAllCards()
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}

	message := fmt.Sprintf("Anzahl der Karten: %d", len(*cards))
	logrus.Debug(message)

	ctx.JSONP(200, message)
}

func (crh *CardRetrieveHandler) GetCard(ctx *gin.Context) {
	queryID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
	}

	logrus.Debugf("API-Handler -> Retrieve card [%d]", queryID)

	card, err := crh.YGOClient.GetCard(queryID)
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}

	logrus.Debug(card)

	ctx.JSONP(200, card)
}

func (crh *CardRetrieveHandler) GetRandomCard(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve random card")

	cards, err := crh.YGOClient.GetAllCards()
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}
	cardsBox := *cards

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randomCard := cardsBox[rand.Intn(len(cardsBox))]

	logrus.Debug(randomCard)

	ctx.JSONP(200, randomCard)
}
