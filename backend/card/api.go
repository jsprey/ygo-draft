package card

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
	"ygodraft/backend/ygoprodeck"
)

type ygoClient interface {
	GetAllCards() (*[]model.Card, error)
}

func SetupAPI(router *gin.RouterGroup, _ *config.YGOContext) {
	cardClient := ygoprodeck.NewYgoProDeckClient()
	cardRetriever := CardRetrieveHandler{
		YGOClient: cardClient,
	}

	router.GET("cards", cardRetriever.GetCards)
}

type CardRetrieveHandler struct {
	YGOClient ygoClient
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
