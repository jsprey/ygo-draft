package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
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

func (ah *authenticationHandler) Login(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Sign in user")

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	// type contains username and password
	loginCredentials := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err = json.Unmarshal(body, loginCredentials)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read crendentials from body: %w", err))
		return
	}

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

	loginResponse := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	ctx.JSON(http.StatusOK, loginResponse)
}

func (ah *authenticationHandler) SignOut(_ *gin.Context) {
	logrus.Debugf("API-Handler -> Sign out user")
}

func (ah *authenticationHandler) Signup(_ *gin.Context) {
	logrus.Debugf("API-Handler -> Sign up user")
}
