package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ygodraft/backend/client/cache"
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

func (crh *CardRetrieveHandler) checkSetExists(ctx *gin.Context, setCode string) (*model.CardSet, error) {
	set, err := crh.YGOClient.GetSet(setCode)
	if err != nil && errors.Is(err, cache.ErrorSetDoesNotExist) {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	} else if err != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, err)
		return nil, err
	}

	return set, err
}

func (crh *CardRetrieveHandler) GetSet(ctx *gin.Context) {
	setCode := ctx.Param("code")
	logrus.Debugf("API-Handler -> Retrieve set with code %s", setCode)

	set, err := crh.checkSetExists(ctx, setCode)
	if err != nil {
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

func (crh *CardRetrieveHandler) GetSetCards(ctx *gin.Context) {
	setCode := ctx.Param("code")
	logrus.Debugf("API-Handler -> Retrieve cards from set with code %s", setCode)

	set, err := crh.checkSetExists(ctx, setCode)
	if err != nil {
		return
	}

	setFilter := model.CardFilter{Sets: []string{set.SetName}}
	cards, err := crh.YGOClient.GetAllCardsWithFilter(setFilter)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if *cards != nil {
		getSetCardsResponse := struct {
			Set   *model.CardSet `json:"set"`
			Cards []*model.Card  `json:"cards"`
		}{Set: set, Cards: *cards}

		ctx.JSONP(http.StatusOK, getSetCardsResponse)
	} else {
		ctx.JSONP(http.StatusNotFound, err)
	}
}
