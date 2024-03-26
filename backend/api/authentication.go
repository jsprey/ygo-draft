package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"ygodraft/backend/model"
)

type authenticationHandler struct {
	ygoAuthClient model.YGOJwtAuthClient
	usermgtClient model.UsermgtClient
}

func newAuthenticationHandler(ygoAuthClient model.YGOJwtAuthClient, usermgtClient model.UsermgtClient) *authenticationHandler {
	return &authenticationHandler{ygoAuthClient: ygoAuthClient, usermgtClient: usermgtClient}
}

// Login Endpoint used to authenticate users.
// @Summary The user provides his credentials and receives a valid JWT that can be used for authentication against the backend server.
// @Description The user provides his credentials and receives a valid JWT that can be used for authentication against the backend server.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body api.Login.LoginCredentials true "No further comment"
// @Failure 401
// @Success 200 {object} api.Login.LoginResponse
// @Router /login [post]
func (ah *authenticationHandler) Login(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Login -> Endpoint called...")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	// type contains username and password
	type LoginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	loginCredentials := &LoginCredentials{}

	err = json.Unmarshal(body, loginCredentials)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read crendentials from body: %w", err))
		return
	}
	logrus.Debugf("API-Handler -> Login -> Login request for user [%s]...", asteriskEmail(loginCredentials.Email))

	// check user exists in database
	user, err := ah.usermgtClient.GetUser(loginCredentials.Email)
	if model.IsErrorUserDoesNotExist(err) {
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to get user: %w", err))
		return
	} else if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get user: %w", err))
		return
	}

	// check password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginCredentials.Password))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to compare passwords: %w", err))
		return
	}

	// create new token with sufficient permissions
	token, err := ah.ygoAuthClient.GenerateToken(*user)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create user token: %w", err))
		return
	}

	type LoginResponse struct {
		Token string `json:"token"`
	}
	loginResponse := &LoginResponse{
		Token: token,
	}

	ctx.JSON(http.StatusOK, loginResponse)
	logrus.Debugf("API-Handler -> Login -> Login granted for user [%s]...", asteriskEmail(loginCredentials.Email))
}
