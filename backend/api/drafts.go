package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/client/draft"
	"ygodraft/backend/customerrors"
	"ygodraft/backend/model"
)

const IDParameter = "id"

type draftsHandler struct {
	DraftClient       model.DraftClient
	DraftPointsClient model.DraftPointsClient
	DraftStoreClient  model.DraftStoreClient
	UsermgtClient     model.UsermgtClient
}

func newDraftsHandler(dbClient model.DatabaseClient, usermgtClient model.UsermgtClient) (*draftsHandler, error) {
	draftClient, err := draft.NewDraftClient(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new draft client: %w", err)
	}

	draftPointsClient, err := draft.NewDraftPointsClient(dbClient, draftClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new draft client: %w", err)
	}

	draftStoreClient, err := draft.NewDraftStoreClient(dbClient, draftClient, draftPointsClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new draft client: %w", err)
	}

	return &draftsHandler{
		UsermgtClient:     usermgtClient,
		DraftPointsClient: draftPointsClient,
		DraftStoreClient:  draftStoreClient,
		DraftClient:       draftClient,
	}, nil
}

// GetDraft Endpoint used to get a specific draft.
// @Summary  Get a specific draft.
// @Description Get a specific draft.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Success 200 {object} api.GetDrafts.getDraftsResponse
// @Failure 400 {string} string "Missing draft id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Draft not found."
// @Failure 404 {string} string "No access to draft."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts [get]
func (dh *draftsHandler) GetDraft(ctx *gin.Context) {
	draftID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a draft id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target draft id: %w", err))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	currentDraft, err := dh.DraftClient.GetDraft(draftID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft by the given id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "Draft not found.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.JSON(http.StatusOK, &currentDraft)
}

// GetDraftChallenges Endpoint used to retrieve all challenges (outgoing + incoming) from the current user.
// @Summary Retrieve all draft challenges (outgoing + incoming) from the current user.
// @Description Retrieve all draft challenges (outgoing + incoming) from the current user.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Success 200 {object} api.GetDraftChallenges.getDraftChallengesResponse
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal server error. Check server logs for more information."
// @Router /drafts/challenges [get]
func (dh *draftsHandler) GetDraftChallenges(ctx *gin.Context) {
	type getDraftChallengesResponse struct {
		Drafts []model.Draft `json:"drafts"`
	}

	logrus.Debugf("API-Handler -> Call to GetDraftChallenges endoint...")

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	challengeDrafts, err := dh.DraftClient.GetDraftsWithStatus(tokenClaims.ID, model.DraftStatusPending)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.JSON(http.StatusOK, &getDraftChallengesResponse{
		Drafts: challengeDrafts,
	})
}

// CreateDraftChallenge Endpoint used to challenge a friend to a draft.
// @Summary Challenge a friend to a draft.
// @Description Challenge a friend to a draft.
// @Tags Draft
// @Security Bearer
// @Accept json
// @Produce json
// @Param receiver body api.CreateDraftChallenge.createDraftChallengeRequest true "Contains the information for the receiving party."
// @Success 204
// @Failure 400 {string} string "Body data is not correct/valid"
// @Failure 401 {string} string "Unauthorized."
// @Failure 409 {string} string "You already have a pending draft challenge to the receiving user."
// @Failure 409 {string} string "You are currently in a running draft against the receiving user and cannot challenge him, until you resolve the draft."
// @Failure 500 {string} string "Internal server error. Check server logs for more information."
// @Router /drafts [post]
func (dh *draftsHandler) CreateDraftChallenge(ctx *gin.Context) {
	type createDraftChallengeRequest struct {
		FriendID int                 `json:"friend_id"`
		Settings model.DraftSettings `json:"settings"`
	}

	logrus.Debugf("API-Handler -> Call to GetDraftChallenges endoint...")

	requestData := &createDraftChallengeRequest{}
	err := GetRequestData(ctx, requestData)
	if err != nil {
		ctx.String(http.StatusBadRequest, "your provided request data is not valid")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	err = dh.DraftClient.CreateDraftChallenge(tokenClaims.ID, requestData.FriendID, requestData.Settings)
	if model.IsErrorCustom(err, model.ErrorUserAlreadyChallenged) {
		ctx.String(http.StatusConflict, fmt.Sprintf("You already have a pending draft challenge to the receiving user."))
		_ = ctx.AbortWithError(http.StatusConflict, customerrors.GenericError(err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserAlreadyRunningDraft) {
		ctx.String(http.StatusConflict, fmt.Sprintf("You are currently in a running draft against the receiving user and cannot challenge him, until you resolve the draft."))
		_ = ctx.AbortWithError(http.StatusConflict, customerrors.GenericError(err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// AcceptDraftChallenge Endpoint used to accept a draft challenge from another user.
// @Summary Accept a draft challenge from another user.
// @Description Accept a draft challenge from another user.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Param id path int true "Contains the id of the challenge to be accepted."
// @Success 204
// @Failure 400 {string} string "You need to provide a draft id."
// @Failure 400 {string} string "The draft is not a challenge and can neither be accepted or declined."
// @Failure 401 {string} string "Unauthorized."
// @Failure 403 {string} string "Only the receiving party can accept the challenge."
// @Failure 404 {string} string "There is no draft by the given id."
// @Failure 409 {string} string "Only the receiving party can accept the challenge."
// @Failure 500 {string} string "Internal server error. Check server logs for more information."
// @Router /drafts/{id}/accept [post]
func (dh *draftsHandler) AcceptDraftChallenge(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Call to AcceptDraftChallenge endoint...")

	claims, currentDraft, err := extractRequestingParty(ctx, dh)
	if err != nil {
		return
	}

	err = dh.DraftClient.AcceptDraftChallenge(currentDraft.ID, claims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftIsNotAChallenge) {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("The draft is not a challenge and can neither be accepted or declined."))
		_ = ctx.AbortWithError(http.StatusBadRequest, customerrors.GenericError(err))
		return
	} else if model.IsErrorCustom(err, model.ErrorOnlyReceivingPartyCanAcceptChallenge) {
		ctx.String(http.StatusForbidden, fmt.Sprintf("Only the receiving party can accept the challenge."))
		_ = ctx.AbortWithError(http.StatusForbidden, customerrors.GenericError(err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DeclineChallenge Endpoint used to decline a draft challenge from another user.
// @Summary Decline a draft challenge from another user.
// @Description Decline a draft challenge from another user.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Param id path int true "Contains the id of the challenge to be declined."
// @Success 204
// @Failure 400 {string} string "You need to provide a draft id."
// @Failure 400 {string} string "The draft is not a challenge and can neither be accepted or declined."
// @Failure 401 {string} string "Unauthorized."
// @Failure 403 {string} string "Only receiving party can decline the challenge."
// @Failure 404 {string} string "There is no draft by the given id."
// @Failure 500 {string} string "Internal server error. Check server logs for more information."
// @Router /drafts/{id}/decline [post]
func (dh *draftsHandler) DeclineChallenge(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Call to DeclineChallenge endoint...")

	claims, currentDraft, err := extractRequestingParty(ctx, dh)
	if err != nil {
		return
	}

	err = dh.DraftClient.DeclineDraftChallenge(currentDraft.ID, claims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftIsNotAChallenge) {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("The draft is not a challenge and can neither be accepted or declined."))
		_ = ctx.AbortWithError(http.StatusBadRequest, customerrors.GenericError(err))
		return
	} else if model.IsErrorCustom(err, model.ErrorOnlyReceivingPartyCanDeclineChallenge) {
		ctx.String(http.StatusForbidden, fmt.Sprintf("Only receiving party can decline the challenge."))
		_ = ctx.AbortWithError(http.StatusForbidden, customerrors.GenericError(err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// PostSurrenderDraft Endpoint used to surrender a certain draft.
// @Summary  Surrender a draft.
// @Description Surrender a draft.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Param id path int true "Contains the id of the draft round."
// @Success 204
// @Failure 400 {string} string "You need to provide a draft id."
// @Failure 400 {string} string "The draft is not a challenge and can neither be accepted or declined."
// @Failure 401 {string} string "Unauthorized."
// @Failure 403 {string} string "Only the receiving party can accept the challenge."
// @Failure 404 {string} string "There is no draft by the given id."
// @Failure 409 {string} string "Only the receiving party can accept the challenge."
// @Failure 500 {string} string "Internal server error. Check server logs for more information."
// @Router /drafts/{id}/surrender [post]
func (dh *draftsHandler) PostSurrenderDraft(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Call to PostSurrenderDraft endoint...")

	draftID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a draft id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target user id: %w", err))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	err = dh.DraftClient.SurrenderRunningDraft(draftID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft by the given id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "You are not part of this draft, and, thus, cannot surrender.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorDraftIsNotRunning) {
		ctx.String(http.StatusNotFound, "The provided draft is not running, and, thus, cannot be surrendered.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}
}

func extractRequestingParty(ctx *gin.Context, dh *draftsHandler) (*model.YgoClaims, model.Draft, error) {
	draftID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a draft id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target user id: %w", err))
		return nil, model.Draft{}, err
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return nil, model.Draft{}, err
	}

	currentDraft, err := dh.DraftClient.GetDraft(draftID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft by the given id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return nil, model.Draft{}, err
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "You are not part of this draft.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return nil, model.Draft{}, err
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return nil, model.Draft{}, err
	}

	return tokenClaims, currentDraft, nil
}

// GetDrafts Endpoint used to get all the running drafts for the current user.
// @Summary  Get all the running drafts for the current user.
// @Description Get all the running drafts for the current user.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Success 200 {object} api.GetDrafts.getDraftsResponse
// @Failure 400 {string} string "Body data is not correct/valid"
// @Failure 401 {string} string "Unauthorized."
// @Failure 403 {string} string "Thrown when anyone except the receiver tries to decline a challenge."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts [get]
func (dh *draftsHandler) GetDrafts(ctx *gin.Context) {
	type getDraftsResponse struct {
		Drafts []model.Draft `json:"drafts"`
	}

	logrus.Debugf("API-Handler -> Call to GetDrafts endoint...")

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	runningDrafts, err := dh.DraftClient.GetDraftsWithStatus(tokenClaims.ID, model.DraftStatusRunning)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.JSON(http.StatusOK, &getDraftsResponse{
		Drafts: runningDrafts,
	})
}

// GetDraftRounds Endpoint used to get all the draft rounds of a given draft.
// @Summary  Get all the draft rounds of a given draft.
// @Description Get all the draft rounds of a given draft.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Param id path int true "Contains the id of the draft to acquire the rounds from."
// @Success 200 {object} api.GetDraftRounds.getDraftRoundsResponse
// @Failure 400 {string} string "Missing draft id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Draft not found."
// @Failure 404 {string} string "No access to draft."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts/{id}/rounds [get]
func (dh *draftsHandler) GetDraftRounds(ctx *gin.Context) {
	type getDraftRoundsResponse struct {
		Rounds []model.DraftRound `json:"rounds"`
	}

	logrus.Debugf("API-Handler -> Call to GetDraftRounds endoint...")

	draftID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a draft id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target draft id: %w", err))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	// check access to draft
	_, err = dh.DraftClient.GetDraft(draftID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft by the given id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "Draft not found.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft with id [%d]: %w", draftID, err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	draftRounds, err := dh.DraftClient.GetDraftRounds(draftID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.JSON(http.StatusOK, getDraftRoundsResponse{Rounds: draftRounds})
}

// GetDraftRoundDecks Endpoint used to get both decks for the draft rounds.
// @Summary  Get both decks for the draft rounds.
// @Description Get both decks for the draft rounds.
// @Tags Draft
// @Security Bearer
// @Produce json
// @Param id path int true "Contains the id of the draft round."
// @Success 200 {object} api.GetDraftRoundDecks.getDraftRoundDecksResponse
// @Failure 400 {string} string "Missing round id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Round not found."
// @Failure 404 {string} string "No access to draft."
// @Failure 500 {string} string "Internal Server Error."
// @Router /rounds/{id}/decks [get]
func (dh *draftsHandler) GetDraftRoundDecks(ctx *gin.Context) {
	type getDraftRoundDecksResponse struct {
		UserDeck  []string `json:"user_deck"`
		EnemyDeck []string `json:"enemy_deck"`
	}

	logrus.Debugf("API-Handler -> Call to GetDraftRoundDecks endoint...")

	roundID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a round id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target round id: %w", err))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "Unauthorized.")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	draftRound, err := dh.DraftClient.GetDraftRound(roundID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft associated with the provided round id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorDraftRoundDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft round by the given id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "Draft not found.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	currentDraft, err := dh.DraftClient.GetDraft(draftRound.DraftID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft associated with the provided round id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "Draft not found.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	userDeck, err := dh.DraftClient.GetDraftRoundDeck(draftRound.ID, tokenClaims.ID)
	if model.IsErrorCustom(err, model.ErrorDraftRoundDeckDoesNotExist) {
		userDeck = []string{}
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	enemyUserID := currentDraft.ReceiverID
	if currentDraft.ReceiverID == tokenClaims.ID {
		enemyUserID = currentDraft.ChallengerID
	}
	enemyDeck, err := dh.DraftClient.GetDraftRoundDeck(draftRound.ID, enemyUserID)
	if model.IsErrorCustom(err, model.ErrorDraftRoundDeckDoesNotExist) {
		enemyDeck = []string{}
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	response := &getDraftRoundDecksResponse{}
	response.UserDeck = userDeck
	response.EnemyDeck = enemyDeck

	ctx.JSON(http.StatusOK, response)

}

// PostDraftRoundDecks Endpoint used to submit a deck for a deck round.
// @Summary  Submit a deck for a deck round.
// @Description Submit a deck for a deck round.
// @Tags Draft
// @Security Bearer
// @Accept json
// @Produce json
// @Param request body api.PostDraftRoundDecks.postDraftRoundDeckRequest true "Contains the id of the draft round."
// @Param id path int true "Contains the id of the draft round."
// @Success 200 {object} api.GetDraftRoundDecks.getDraftRoundDecksResponse
// @Failure 400 {string} string "Missing round id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Round not found."
// @Failure 404 {string} string "No access to draft."
// @Failure 500 {string} string "Internal Server Error."
// @Router /rounds/{id}/decks [post]
func (dh *draftsHandler) PostDraftRoundDecks(ctx *gin.Context) {
	type postDraftRoundDeckRequest struct {
		Deck []string `json:"deck"`
	}

	logrus.Debugf("API-Handler -> Call to PostDraftRoundDecks endoint...")

	roundID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a round id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target round id: %w", err))
		return
	}

	requestData := &postDraftRoundDeckRequest{}
	err = GetRequestData(ctx, requestData)
	if err != nil {
		ctx.String(http.StatusBadRequest, "your provided request data is not valid")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	userDeck, err := dh.DraftClient.GetDraftRoundDeck(roundID, tokenClaims.ID)
	if err != nil && !model.IsErrorCustom(err, model.ErrorDraftRoundDeckDoesNotExist) {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	if len(userDeck) > 0 {
		ctx.String(http.StatusConflict, "A deck was already submitted!")
		_ = ctx.AbortWithError(http.StatusConflict, customerrors.GenericError(err))
		return
	}

	err = dh.DraftClient.SubmitDraftDeck(roundID, tokenClaims.ID, requestData.Deck)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// PostRoundWinner Endpoint used to set the winner of a draft round.
// @Summary  Set the winner of a draft round.
// @Description Set the winner of a draft round.
// @Tags Draft
// @Security Bearer
// @Accept json
// @Produce json
// @Param request body api.PostRoundWinner.postRoundWinnerRequest true "Contains the id of the draft round."
// @Param id path int true "Contains the id of the draft round."
// @Success 204
// @Failure 400 {string} string "Missing round id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Round not found."
// @Failure 404 {string} string "No access to draft."
// @Failure 500 {string} string "Internal Server Error."
// @Router /rounds/{id} [post]
func (dh *draftsHandler) PostRoundWinner(ctx *gin.Context) {
	type postRoundWinnerRequest struct {
		Winner int `json:"winner"`
	}

	logrus.Debugf("API-Handler -> Call to PostRoundWinner endoint...")

	roundID, err := strconv.Atoi(ctx.Param(IDParameter))
	if err != nil {
		ctx.String(http.StatusBadRequest, "You need to provide a round id.")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target round id: %w", err))
		return
	}

	requestData := &postRoundWinnerRequest{}
	err = GetRequestData(ctx, requestData)
	if err != nil {
		ctx.String(http.StatusBadRequest, "your provided request data is not valid")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	err = dh.DraftClient.SetWinnerForDraftRound(roundID, requestData.Winner)
	if model.IsErrorCustom(err, model.ErrorDraftDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft associated with the provided round id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorDraftRoundDoesNotExist) {
		ctx.String(http.StatusNotFound, "There is no draft round by the given id.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if model.IsErrorCustom(err, model.ErrorUserIsNotParticipatingInDraft) {
		ctx.String(http.StatusNotFound, "Draft not found.")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get draft round with id [%d]: %w", roundID, err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
