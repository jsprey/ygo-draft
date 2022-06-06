package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
)

func (crh *CardRetrieveHandler) GetSets(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve all sets")

	sets, err := crh.YGOClient.GetAllSets()
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}

	if *sets != nil {
		randomCardsResponse := struct {
			Sets []*model.CardSet `json:"sets"`
		}{Sets: *sets}

		ctx.JSONP(200, randomCardsResponse)
	} else {
		emptyResponse := struct {
			Sets []*model.CardSet `json:"sets"`
		}{Sets: []*model.CardSet{}}

		ctx.JSONP(200, emptyResponse)
	}
}
