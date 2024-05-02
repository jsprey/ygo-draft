package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/customerrors"
	"ygodraft/backend/model"
)

// GetDraftShopBalance Endpoint used to get the points of a current user in a specific draft.
// @Summary  Get the points of a current user in a specific draft.
// @Description Get the points of a current user in a specific draft.
// @Tags Draft - Shop
// @Security Bearer
// @Produce json
// @Param id path int true "Contains the id of the draft."
// @Success 200 {object} api.GetDraftShopBalance.getDraftPointsResponse
// @Failure 400 {string} string "Missing draft id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Draft not found."
// @Failure 404 {string} string "No access to draft."
// @Failure 500 {string} string "Internal Server Error."
// @Router /drafts/{id}/shop/balance [get]
func (dh *draftsHandler) GetDraftShopBalance(ctx *gin.Context) {
	type getDraftPointsResponse struct {
		Points int `json:"points"`
	}

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

	currentPoints, err := dh.DraftPointsClient.GetPoints(draftID, tokenClaims.ID)
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

	ctx.JSON(http.StatusOK, &getDraftPointsResponse{Points: currentPoints.Points})
}
