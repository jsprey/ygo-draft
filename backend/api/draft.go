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

const AcceptChallengeParamID = "id"

type draftHandler struct {
	ChallengeClient model.DraftChallengeClient
	DraftClient     model.DraftClient
	UsermgtClient   model.UsermgtClient
}

func newDraftHandler(dbClient model.DatabaseClient, usermgtClient model.UsermgtClient) (*draftHandler, error) {
	challengeClient, err := draft.NewChallengeClient(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new challenge client: %w", err)
	}

	draftClient, err := draft.NewDraftClient(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new draft client: %w", err)
	}

	return &draftHandler{
		ChallengeClient: challengeClient,
		DraftClient:     draftClient,
		UsermgtClient:   usermgtClient,
	}, nil
}

// GetChallenges Endpoint used to retrieve the challenges of the current user.
// @Summary Retrieve the challenges of the current user.
// @Description Retrieve the challenges of the current user.
// @Tags Draft
// @Security ApiKeyAuth
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Success 200 {object} api.GetChallenges.getChallengesResponse
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts/challenges [get]
func (dh *draftHandler) GetChallenges(ctx *gin.Context) {
	type getChallengesResponse struct {
		Challenges []model.DraftChallenge `json:"challenges"`
	}

	logrus.Debugf("API-Handler -> Call to GetChallenges endoint...")

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	challenges, err := dh.ChallengeClient.GetChallenges(tokenClaims.ID, model.StatusPending)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.JSON(http.StatusOK, &getChallengesResponse{
		Challenges: challenges,
	})
}

// ChallengeFriend Endpoint used to challenge a friend to a draft.
// @Summary Challenge a friend to a draft.
// @Description Challenge a friend to a draft.
// @Tags Draft
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Param receiver body api.ChallengeFriend.challengeFriendRequest true "Contains the information for the receiving party."
// @Success 204
// @Failure 400 {string} string "Body data is not correct/valid"
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts/challenges [post]
func (dh *draftHandler) ChallengeFriend(ctx *gin.Context) {
	type challengeFriendRequest struct {
		FriendID int                 `json:"friend_id"`
		Settings model.DraftSettings `json:"settings"`
	}

	logrus.Debugf("API-Handler -> Call to GetChallenges endoint...")

	requestData := &challengeFriendRequest{}
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

	err = dh.ChallengeClient.ChallengeUser(tokenClaims.ID, requestData.FriendID, requestData.Settings)
	if err != nil && !model.IsErrorUserAlreadyChallenged(err) {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// AcceptChallenge Endpoint used to accept a draft challenge from another user.
// @Summary Accept a draft challenge from another user.
// @Description Accept a draft challenge from another user.
// @Tags Draft
// @Security ApiKeyAuth
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Param id path int true "Contains the id of the challenge to be accepted."
// @Success 204
// @Failure 400 {string} string "Body data is not correct/valid"
// @Failure 401 {string} string "Unauthorized."
// @Failure 403 {string} string "Thrown when anyone except the receiver tries to accept a challenge."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts/challenges/{id}/accept [post]
func (dh *draftHandler) AcceptChallenge(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Call to AcceptChallenge endoint...")

	challenge, err := extractChallengeReceiver(ctx, dh)
	if err != nil {
		return
	}

	if challenge.Status == model.StatusAccepted {
		ctx.Status(http.StatusNoContent)
		return
	}

	if challenge.Status == model.StatusDeclined {
		ctx.String(http.StatusBadRequest, "challenges was already declined")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("declined challenge cannot be accepted: %w", err))
		return
	}

	err = dh.ChallengeClient.AcceptChallenge(challenge.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func extractChallengeReceiver(ctx *gin.Context, dh *draftHandler) (model.DraftChallenge, error) {
	challengeID, err := strconv.Atoi(ctx.Param(AcceptChallengeParamID))
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid challenge id")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read the target user id: %w", err))
		return model.DraftChallenge{}, err
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return model.DraftChallenge{}, err
	}

	challenge, err := dh.ChallengeClient.GetChallenge(challengeID)
	if model.IsErrorChallengeDoesNotExist(err) {
		ctx.String(http.StatusBadRequest, "you need to provide a valid challenge id")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to get challenge with id [%d]: %w", challengeID, err))
		return model.DraftChallenge{}, err
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return model.DraftChallenge{}, err
	}

	if challenge.ReceiverID != tokenClaims.ID {
		ctx.String(http.StatusForbidden, "you are not the receiver of this challenge")
		_ = ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("non receiving user tries to accept a challenge: %w", err))
		return model.DraftChallenge{}, err
	}

	return challenge, nil
}

// DeclineChallenge Endpoint used to decline a draft challenge from another user.
// @Summary Decline a draft challenge from another user.
// @Description Decline a draft challenge from another user.
// @Tags Draft
// @Security ApiKeyAuth
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Param id path int true "Contains the id of the challenge to be declined."
// @Success 204
// @Failure 400 {string} string "Body data is not correct/valid"
// @Failure 401 {string} string "Unauthorized."
// @Failure 403 {string} string "Thrown when anyone except the receiver tries to decline a challenge."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts/challenges/{id}/decline [post]
func (dh *draftHandler) DeclineChallenge(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Call to DeclineChallenge endoint...")

	challenge, err := extractChallengeReceiver(ctx, dh)
	if err != nil {
		return
	}

	if challenge.Status == model.StatusDeclined {
		ctx.Status(http.StatusNoContent)
		return
	}

	if challenge.Status == model.StatusAccepted {
		ctx.String(http.StatusBadRequest, "challenges was already accepted")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("accepted challenge cannot be declined: %w", err))
		return
	}

	err = dh.ChallengeClient.DeclineChallenge(challenge.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, customerrors.GenericError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
