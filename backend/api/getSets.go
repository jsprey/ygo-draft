package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ygodraft/backend/model"
)

func (crh *CardRetrieveHandler) GetSets(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve all sets")

	sets, err := crh.YGOClient.GetAllSets()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if *sets != nil {
		randomCardsResponse := struct {
			Sets []*model.CardSet `json:"sets"`
		}{Sets: *sets}

		ctx.JSONP(http.StatusOK, randomCardsResponse)
	} else {
		emptyResponse := struct {
			Sets []*model.CardSet `json:"sets"`
		}{Sets: []*model.CardSet{}}

		ctx.JSONP(http.StatusOK, emptyResponse)
	}
}

func (crh *CardRetrieveHandler) GetSet(ctx *gin.Context) {
	setCode := ctx.Param("code")
	logrus.Debugf("API-Handler -> Retrieve set with code %s", setCode)

	set, err := crh.YGOClient.GetSet(setCode)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if set != nil {
		setResponse := struct {
			Set *model.CardSet `json:"set"`
		}{Set: set}

		ctx.JSONP(http.StatusOK, setResponse)
	} else {
		ctx.JSONP(http.StatusNotFound, err)
	}
}
