package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"net/mail"
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
	logrus.Debugf("API-Handler -> Login -> Endpoint called...")

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

	loginResponse := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	ctx.JSON(http.StatusOK, loginResponse)
	logrus.Debugf("API-Handler -> Login -> Login granted for user [%s]...", asteriskEmail(loginCredentials.Email))
}

func (ah *authenticationHandler) DeleteUser(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> DeleteUser -> Call to DeleteUser endoint...")

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	// type contains username and password
	deleteRequest := &struct {
		Email string `json:"email"`
	}{}

	err = json.Unmarshal(body, deleteRequest)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read delete request from body: %w", err))
		return
	}
	logrus.Debugf("API-Handler -> DeleteUser -> Delete request for user [%s]...", asteriskEmail(deleteRequest.Email))

	// check user exists in database
	user, err := ah.usermgtClient.GetUser(deleteRequest.Email)
	if model.IsErrorUserDoesNotExist(err) {
		ctx.String(http.StatusNotFound, "user does not exist")
		_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get suer: %w", err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get user: %w", err))
		return
	}

	err = ah.usermgtClient.DeleteUser(user.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to delete user: %w", err))
		return
	}

	ctx.Status(204)
	logrus.Debugf("API-Handler -> DeleteUser -> User [%s] successfully deleted.", asteriskEmail(deleteRequest.Email))
}

func (ah *authenticationHandler) RegisterUser(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> RegisterUser -> Call to RegisterUser endoint...")

	// read request data
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid user registration request")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	userRegistrationRequest := &struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		DisplayName string `json:"display_name"`
		IsAdmin     bool   `json:"is_admin"`
	}{}
	err = json.Unmarshal(body, userRegistrationRequest)
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid user registration request")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to unmarshal body into user registration request: %w", err))
	}
	logrus.Debugf("API-Handler -> RegisterUser -> Request to register user [%s]...", asteriskEmail(userRegistrationRequest.Email))

	// validate email exists
	if userRegistrationRequest.Email == "" {
		ctx.String(http.StatusBadRequest, "the field 'email' is missing or empty")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("the field 'email' is missing or empty"))
		return
	}

	// validate email is valid
	_, err = mail.ParseAddress(userRegistrationRequest.Email)
	if err != nil {
		ctx.String(http.StatusBadRequest, "provided email is not a valid email address")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to validate email address: %w", err))
		return
	}

	// validate password exists
	if userRegistrationRequest.Password == "" {
		ctx.String(http.StatusBadRequest, "the field 'password' is missing or empty")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("the field 'password' is missing or empty"))
		return
	}

	// validate password is long enough
	if len(userRegistrationRequest.Password) < 8 {
		ctx.String(http.StatusBadRequest, "password must be at least 8 characters long")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("password must be at least 8 characters long"))
		return
	}

	// check user exists in database
	user, err := ah.usermgtClient.GetUser(userRegistrationRequest.Email)
	if err != nil && !model.IsErrorUserDoesNotExist(err) {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get user: %w", err))
		return
	} else if err == nil && user != nil {
		ctx.String(http.StatusConflict, "user already exists")
		_ = ctx.AbortWithError(http.StatusConflict, fmt.Errorf("user already exists"))
		return
	}

	// hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistrationRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to hash password: %w", err))
		return
	}

	// create new user in database
	newUser := model.User{
		Email:        userRegistrationRequest.Email,
		PasswordHash: string(hashedPassword),
		DisplayName:  userRegistrationRequest.DisplayName,
		IsAdmin:      userRegistrationRequest.IsAdmin,
	}
	err = ah.usermgtClient.CreateUser(newUser)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create user: %w", err))
	}

	ctx.Status(204)
	logrus.Debugf("API-Handler -> RegisterUser -> Registered the user [%s] successfully.", asteriskEmail(newUser.Email))
}

func (ah *authenticationHandler) CurrentUser(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> CurrentUser -> Call to CurrentUser endoint...")

	token, err := ExtractBearerToken(ctx)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "you need to provide a valid bearer token")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to extract bearer token: %w", err))
	}

	tokenClaims, err := ah.ygoAuthClient.ValidateToken(token)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "your provided token is not valid")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to validate token: %w", err))
	}

	user, err := ah.usermgtClient.GetUser(tokenClaims.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get user: %w", err))
		return
	}

	userReponse := &struct {
		ID          int    `json:"id"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
		IsAdmin     bool   `json:"is_admin"`
	}{
		ID:          user.ID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		IsAdmin:     user.IsAdmin,
	}
	ctx.JSON(http.StatusOK, userReponse)
}

func asteriskEmail(email string) string {
	asteriskEmail := string([]rune(email)[:5]) + "**********"
	return asteriskEmail
}
