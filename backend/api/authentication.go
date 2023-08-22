package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"ygodraft/backend/model"
)

type authenticationHandler struct {
	ygoAuthClient model.YGOJwtAuthClient
}

func newAuthenticationHandler(ygoAuthClient model.YGOJwtAuthClient) *authenticationHandler {
	return &authenticationHandler{ygoAuthClient: ygoAuthClient}
}

// NewAuthenticationHandler creates a new authentication handler with the given authentication context.

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

	logrus.Debugf("API-Handler -> Login User [%s] with passwordhash [%s]", loginCredentials.Email, loginCredentials.Password)

	loginResponse := &struct {
		Token string `json:"token"`
	}{
		Token: "myTestToken",
	}

	ctx.JSON(http.StatusInternalServerError, loginResponse)
	// check user exists in database
	// create new token with sufficient permissions
	// return token
}

func (ah *authenticationHandler) SignOut(_ *gin.Context) {
	logrus.Debugf("API-Handler -> Sign out user")
}

func (ah *authenticationHandler) Signup(_ *gin.Context) {
	logrus.Debugf("API-Handler -> Sign up user")
}
