package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
)

const authQueryParamEmail = "email"
const authQueryParamPasswordHash = "password_hash"

type authenticationHandler struct {
	ygoAuthClient model.YGOJwtAuthClient
}

func newAuthenticationHandler(ygoAuthClient model.YGOJwtAuthClient) *authenticationHandler {
	return &authenticationHandler{ygoAuthClient: ygoAuthClient}
}

// NewAuthenticationHandler creates a new authentication handler with the given authentication context.

func (ah *authenticationHandler) Login(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Sign in user")

	email := ctx.Query(authQueryParamEmail)
	pwHash := ctx.Query(authQueryParamPasswordHash)

	logrus.Debugf("API-Handler -> Login User [%s] with passwordhash [%s]", email, pwHash)

	// check user exists in database
	// create new token with sufficient permissions
	// return token
}

func (ah *authenticationHandler) SignOut(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Sign out user")
}

func (ah *authenticationHandler) Signup(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> Sign up user")
}
