package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/mail"
	"strconv"
	"ygodraft/backend/model"
)

const InvalidGivenUserErrorMessage = "Bad request. You need to provide a valid user id."
const BadRequestCannotReferenceYourself = "Bad request. You need to provide another user. Not you own."
const GetUsersPageParameter = "page"
const GetUsersPageSizeParameter = "page_size"

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

func (ah *authenticationHandler) GetFriends(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> CurrentUser -> Call to GetFriends endoint...")

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

	friendList, err := ah.usermgtClient.GetFriends(user.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to retrieve friends of user: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, friendList)
}

func (ah *authenticationHandler) GetFriendRequests(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> CurrentUser -> Call to GetFriendRequests endoint...")

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

	friendRequests, err := ah.usermgtClient.GetFriendRequests(user.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to retrieve friends of user: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, friendRequests)
}

func (ah *authenticationHandler) PostRequestFriend(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> CurrentUser -> Call to PostRequestFriend endoint...")

	targetUserID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid user id")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to read the target user id: %w", err))
	}

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

	targetUser, err := ah.usermgtClient.GetUserByID(targetUserID)
	if err != nil && model.IsErrorUserDoesNotExist(err) {
		ctx.String(http.StatusBadRequest, InvalidGivenUserErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get target user: %w", err))
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get target user: %w", err))
		return
	}

	if user.ID == targetUserID {
		ctx.String(http.StatusBadRequest, BadRequestCannotReferenceYourself)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("invalid operation: %w", err))
		return
	}

	friendRequests, err := ah.usermgtClient.GetFriendRequests(user.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to retrieve friend requests of user: %w", err))
		return
	}

	for _, friendRequest := range friendRequests {
		if friendRequest.ID == targetUser.ID {
			err = ah.usermgtClient.SetRelationshipStatus(user.ID, targetUser.ID, model.FriendStatusFriends)
			if err != nil {
				ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
				_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to set relationship status: %w", err))
				return
			}

			ctx.Status(http.StatusOK)
			return
		}
	}

	// not invited -> send invite to friend
	err = ah.usermgtClient.SetRelationshipStatus(user.ID, targetUser.ID, model.FriendStatusInvited)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to set relationship status: %w", err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (ah *authenticationHandler) PostRequestFriendByEmail(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> CurrentUser -> Call to PostRequestFriendByEmail endoint...")

	rawBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cannot read body of request"))
		return
	}

	var bodyData struct {
		FriendEmail string `json:"friend_email"`
	}
	err = json.Unmarshal(rawBody, &bodyData)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Form data is incorrect/invalid.")
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("bad request missing friend email"))
		return
	}

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

	targetUser, err := ah.usermgtClient.GetUser(bodyData.FriendEmail)
	if err != nil && model.IsErrorUserDoesNotExist(err) {
		logrus.Debug("The user to add does not seem to exist -> skip...")
		ctx.Status(http.StatusOK)
		return
	} else if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get target user by email: %w", err))
		return
	}

	if user.Email == targetUser.Email {
		ctx.String(http.StatusBadRequest, BadRequestCannotReferenceYourself)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("invalid operation: %w", err))
		return
	}

	err = ah.usermgtClient.SetRelationshipStatus(user.ID, targetUser.ID, model.FriendStatusInvited)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to set relationship status: %w", err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (ah *authenticationHandler) GetUsers(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> CurrentUser -> Call to GetUsers endoint...")

	pageParameter, err := GetQueryParameterInt(ctx, GetUsersPageParameter)
	if err != nil {
		ctx.String(http.StatusBadRequest, "%s", err.Error())
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to get query integer parameter [%s]: %w", GetUsersPageParameter, err))
		return
	}

	pageSizeParameter, err := GetQueryParameterInt(ctx, GetUsersPageSizeParameter)
	if err != nil {
		ctx.String(http.StatusBadRequest, "%s", err.Error())
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to get query integer parameter [%s]: %w", GetUsersPageSizeParameter, err))
		return
	}

	usersCount, err := ah.usermgtClient.CountUsers()
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to retrieve number of users: %w", err))
		return
	}

	numberOfPages := int(math.Ceil(float64(usersCount) / float64(pageSizeParameter)))
	if pageParameter >= numberOfPages {
		ctx.String(http.StatusBadRequest, "parameter [%s=%d] cannot exceed number of pages [%d]", GetUsersPageParameter, pageParameter, numberOfPages)
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("parameter [%s=%d] cannot exceed number of pages [%d]", GetUsersPageParameter, pageParameter, numberOfPages))
		return
	}

	userList, err := ah.usermgtClient.GetUsers(pageParameter, pageSizeParameter)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get users: %w", err))
		return
	}

	var getUsersResponse = &struct {
		NumberOfUsers int          `json:"numberOfUsers"`
		NumberOfPages int          `json:"numberOfPages"`
		Users         []model.User `json:"users"`
	}{NumberOfUsers: usersCount, Users: userList, NumberOfPages: numberOfPages}

	ctx.JSON(http.StatusOK, &getUsersResponse)
}
