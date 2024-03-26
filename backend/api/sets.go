package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/model"
)

// GetSets Endpoint used to retrieve all sets.
// @Summary Retrieve all sets.
// @Description Retrieve all sets.
// @Tags Yu-Gi-Oh
// @Produce json
// @Success 200 {object} api.GetSets.getSetsResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /sets [get]
func (crh *ygoRetrieveHandler) GetSets(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve all sets")

	sets, err := crh.YGOClient.GetAllSets()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	type getSetsResponse struct {
		Sets []*model.CardSet `json:"sets"`
	}

	if *sets != nil {
		randomCardsResponse := getSetsResponse{Sets: *sets}
		ctx.JSONP(http.StatusOK, randomCardsResponse)
	} else {
		emptyResponse := getSetsResponse{Sets: []*model.CardSet{}}
		ctx.JSONP(http.StatusOK, emptyResponse)
	}
}

func (crh *ygoRetrieveHandler) checkSetExists(ctx *gin.Context, setCode string) (*model.CardSet, error) {
	set, err := crh.YGOClient.GetSet(setCode)
	if err != nil && errors.Is(err, cache.ErrorSetDoesNotExist) {
		_ = ctx.AbortWithError(http.StatusNotFound, err)
		return nil, err
	} else if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	return set, err
}

// GetSet Endpoint used to retrieve a specific set.
// @Summary Retrieve a specific set.
// @Description Retrieve a specific set.
// @Tags Yu-Gi-Oh
// @Produce json
// @Param code path string true "The identifier of the set."
// @Success 200 {object} api.GetSet.getSetResponse
// @Failure 404 {string} string "Set does not exist."
// @Failure 500 {string} string "Internal Server Error"
// @Router /sets/{code} [get]
func (crh *ygoRetrieveHandler) GetSet(ctx *gin.Context) {
	setCode := ctx.Param("code")
	logrus.Debugf("API-Handler -> Retrieve set with code %s", setCode)

	set, err := crh.checkSetExists(ctx, setCode)
	if err != nil {
		return
	}

	type getSetResponse struct {
		Set *model.CardSet `json:"set"`
	}

	if set != nil {
		setResponse := getSetResponse{Set: set}
		ctx.JSONP(http.StatusOK, setResponse)
	} else {
		ctx.JSONP(http.StatusNotFound, err)
	}
}

// GetSetCards Endpoint used to retrieve all cards of a specific set.
// @Summary Retrieve all cards of a specific set.
// @Description Retrieve all cards of a specific set.
// @Tags Yu-Gi-Oh
// @Produce json
// @Param code path string true "The identifier of the set."
// @Success 200 {object} api.GetSetCards.getSetCardsResponse
// @Failure 404 {string} string "Set does not exist or set contains no cards."
// @Failure 500 {string} string "Internal Server Error"
// @Router /sets/{code}/cards [get]
func (crh *ygoRetrieveHandler) GetSetCards(ctx *gin.Context) {
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

	type getSetCardsResponse struct {
		Set   *model.CardSet `json:"set"`
		Cards []*model.Card  `json:"cards"`
	}

	if *cards != nil {
		response := getSetCardsResponse{Set: set, Cards: *cards}
		ctx.JSONP(http.StatusOK, response)
	} else {
		ctx.JSONP(http.StatusNotFound, err)
	}
}
