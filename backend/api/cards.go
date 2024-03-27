package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"ygodraft/backend/customerrors"
	"ygodraft/backend/model"
)

const GetCardQueryParamID = "id"
const GetRandomCardsQueryParamSize = "size"
const GetRandomCardsQueryParamSets = "sets"
const GetRandomCardsQueryParamTypes = "types"

type ygoRetrieveHandler struct {
	YGOClient model.YgoClient
}

func newYgoRetrieveHandler(client model.YgoClient) *ygoRetrieveHandler {
	return &ygoRetrieveHandler{YGOClient: client}
}

// GetCards Endpoint used to retrieve all cards.
// @Summary Endpoint used to retrieve all cards.
// @Description Endpoint used to retrieve all cards.
// @Tags Yu-Gi-Oh
// @Produce json
// @Success 200 {object} api.GetCards.getCardResponse
// @Failure 500
// @Router /cards [get]
func (crh *ygoRetrieveHandler) GetCards(ctx *gin.Context) {
	cards, err := crh.YGOClient.GetAllCards()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	type getCardResponse struct {
		CardIds []int `json:"card_ids"`
		Number  int   `json:"number"`
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

// GetCard Endpoint used to retrieve a specific card.
// @Summary Endpoint used to retrieve a specific card.
// @Description Endpoint used to retrieve a specific cards.
// @Tags Yu-Gi-Oh
// @Produce json
// @Param cardID path int true "The id of the requested card."
// @Success 200 {object} model.Card
// @Success 400
// @Failure 500
// @Router /cards/{cardID} [get]
func (crh *ygoRetrieveHandler) GetCard(ctx *gin.Context) {
	queryID, err := strconv.Atoi(ctx.Param(GetCardQueryParamID))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Debugf("API-Handler -> Retrieve api [%d]", queryID)

	card, err := crh.YGOClient.GetCard(queryID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	logrus.Debug(card)

	ctx.JSONP(http.StatusOK, card)
}

// GetRandomCards Endpoint used to retrieve a certain amount of random card at once.
// @Summary Retrieve a certain amount of random card at once.
// @Description Retrieve a certain amount of random card at once. Calling the endpoint without any parameters returns
// @Description a only one random card.
// @Tags Yu-Gi-Oh
// @Produce json
// @Param size query int false "Determines the amount of random cards to be returned."
// @Param sets query string false "Contains the filtered sets, separated by comma, e.g., 'Set 1,Set 2,Set 3'"
// @Param types query string false "Contains the filtered types, separated by comma, e.g., 'Fire, Water, Dragon'"
// @Success 200 {object} api.GetRandomCards.getRandomCardsResponse
// @Failure 400
// @Failure 500
// @Router /cards/random [get]
func (crh *ygoRetrieveHandler) GetRandomCards(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Retrieve random cards")

	numberOfCards, cardFilter, err := getRandomCardsCheckQueryAttributes(ctx)
	if err != nil {
		return
	}

	cards, err := crh.YGOClient.GetAllCardsWithFilter(cardFilter)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	type getRandomCardsResponse struct {
		Cards []*model.Card `json:"cards"`
	}

	if *cards != nil {
		cardsBox := *cards

		randomDeck := make([]*model.Card, numberOfCards)
		for i := 0; i < numberOfCards; i++ {
			randomDeck[i] = cardsBox[rand.Intn(len(cardsBox))]
		}
		randomCardsResponse := &getRandomCardsResponse{Cards: randomDeck}

		ctx.JSONP(http.StatusOK, randomCardsResponse)
	} else {
		emptyCards := []*model.Card{}
		emptyResponse := getRandomCardsResponse{Cards: emptyCards}

		ctx.JSONP(http.StatusOK, emptyResponse)
	}
}

func getRandomCardsCheckQueryAttributes(ctx *gin.Context) (int, model.CardFilter, error) {
	filter := model.CardFilter{}
	numberOfCards := 1

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
