package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"math"
	"net/http"
	"net/mail"
	"strconv"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/model"
)

const InvalidGivenUserErrorMessage = "Bad request. You need to provide a valid user id."
const BadRequestCannotReferenceYourself = "Bad request. You need to provide another user. Not you own."
const GetUsersPageParameter = "page"
const GetUsersPageSizeParameter = "page_size"

type userManagementHandler struct {
	usermgtClient model.UsermgtClient
}

func newUserManagementHandler(usermgtClient model.UsermgtClient) *userManagementHandler {
	return &userManagementHandler{usermgtClient: usermgtClient}
}

// GetUsers Endpoint used to retrieve all users with pagination.
// @Summary Retrieve all users with pagination.
// @Description Retrieve all users with pagination.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query int true "Current page for the pagination."
// @Param page_size query int true "The size for the pages."
// @Param authorization header string true "Contains the authorization token."
// @Success 200 {object} api.GetUsers.getUsersResponse
// @Failure 400 {string} string "Missing query parameter."
// @Failure 400 {string} string "The page parameter cannot exceed the available amount of pages."
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [get]
func (ah *userManagementHandler) GetUsers(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> GetCurrentUser -> Call to GetUsers endoint...")

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

	type getUsersResponse = struct {
		NumberOfUsers int          `json:"numberOfUsers"`
		NumberOfPages int          `json:"numberOfPages"`
		Users         []model.User `json:"users"`
	}
	var response = &getUsersResponse{NumberOfUsers: usersCount, Users: userList, NumberOfPages: numberOfPages}

	ctx.JSON(http.StatusOK, response)
}

// DeleteUser Endpoint used to delete a user.
// @Summary Deletes a user.
// @Description Deletes a user.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_data body api.DeleteUser.deleteUserRequest true "Contains the identifier for the to be deleted user."
// @Param authorization header string true "Contains the authorization token."
// @Success 204
// @Failure 401 {string} string "Unauthorized."
// @Failure 400 {string} string "Missing body."
// @Failure 400 {string} string "Body contains invalid data."
// @Failure 404 {string} string "User does not exist."
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [delete]
func (ah *userManagementHandler) DeleteUser(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> DeleteUser -> Call to DeleteUser endoint...")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	// type contains username and password
	type deleteUserRequest struct {
		Email string `json:"email"`
	}
	deleteRequest := &deleteUserRequest{}

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

	err = ah.usermgtClient.DeleteUser(user.ID, user.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to delete user: %w", err))
		return
	}

	ctx.Status(http.StatusNoContent)
	logrus.Debugf("API-Handler -> DeleteUser -> User [%s] successfully deleted.", asteriskEmail(deleteRequest.Email))
}

// PostUsers Endpoint used to create a new user.
// @Summary Creates a new user.
// @Description Creates a new user.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user_data body api.PostUsers.postUsersRequest true "Contains the information for the new user."
// @Param authorization header string true "Contains the authorization token."
// @Success 204
// @Failure 400 {string} string "Missing body."
// @Failure 401 {string} string "Unauthorized."
// @Failure 400 {string} string "Body contains invalid data."
// @Failure 409 {string} string "User already exists."
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
func (ah *userManagementHandler) PostUsers(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> PostUsers -> Call to PostUsers endoint...")

	// read request data
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid user registration request")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to read request body: %w", err))
		return
	}

	log.Printf("%+v", string(body))

	type postUsersRequest = struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		DisplayName string `json:"display_name"`
		IsAdmin     bool   `json:"is_admin"`
	}
	userRegistrationRequest := &postUsersRequest{}
	err = json.Unmarshal(body, userRegistrationRequest)
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid user registration request")
		_ = ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to unmarshal body into user registration request: %w", err))
	}
	logrus.Debugf("API-Handler -> PostUsers -> Request to register user [%s]...", asteriskEmail(userRegistrationRequest.Email))

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
	logrus.Debugf("API-Handler -> PostUsers -> Registered the user [%s] successfully.", asteriskEmail(newUser.Email))
}

// GetCurrentUser Endpoint used to retrieve the current user.
// @Summary Retrieve the current user.
// @Description Retrieve the current user.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Success 200 {object} api.GetCurrentUser.getCurrentUserResponse
// @Failure 400 {string} string "Missing body."
// @Failure 401 {string} string "Unauthorized."
// @Failure 400 {string} string "Body contains invalid data."
// @Failure 409 {string} string "User already exists."
// @Failure 500 {string} string "Internal Server Error"
// @Router /user [get]
func (ah *userManagementHandler) GetCurrentUser(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> GetCurrentUser -> Call to GetCurrentUser endoint...")

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
	}

	user, err := ah.usermgtClient.GetUser(tokenClaims.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get user: %w", err))
		return
	}

	type getCurrentUserResponse = struct {
		ID          int    `json:"id"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
		IsAdmin     bool   `json:"is_admin"`
	}
	userReponse := &getCurrentUserResponse{
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

// GetFriends Endpoint used to retrieve the friends of a user.
// @Summary Retrieve the friends of a user.
// @Description Retrieve the friends of a user.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Success 200 {array} model.Friend
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/friends [get]
func (ah *userManagementHandler) GetFriends(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> GetCurrentUser -> Call to GetFriends endoint...")

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
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

// GetFriendRequests Endpoint used to retrieve the friend requests of a user.
// @Summary Retrieve the friend requests of a user.
// @Description Retrieve the friend requests of a user.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param authorization header string true "Contains the authorization token."
// @Success 200 {array} model.FriendRequest
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/friends/requests [get]
func (ah *userManagementHandler) GetFriendRequests(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> GetCurrentUser -> Call to GetFriendRequests endoint...")

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
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

// PostFriendRequest Endpoint used to send/accept a friend request from the current user to another.
// @Summary send/accept a friend request from the current user to another.
// @Description send/accept a friend request from the current user to another. When the target user does not exist, nothing happens.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param targetUser path int true "Contains the id of the target user."
// @Param authorization header string true "Contains the authorization token."
// @Success 201
// @Failure 400 {string} string "Cannot post a request to yourself."
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/friends/requests/{targetUser} [post]
func (ah *userManagementHandler) PostFriendRequest(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> GetCurrentUser -> Call to PostFriendRequest endoint...")

	targetUserID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "you need to provide a valid user id")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to read the target user id: %w", err))
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
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

	ctx.Status(http.StatusNoContent)
}

// PostFriendRequestByEmail Endpoint used to send/accept a friend request from the current user to another via email.
// @Summary send/accept a friend request from the current user to another.
// @Description send/accept a friend request from the current user to another. When the target user does not exist, nothing happens.
// @Tags User Management
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param targetUser body int true "Contains the email of the target user."
// @Param authorization header string true "Contains the authorization token."
// @Success 204
// @Failure 400 {string} string "Cannot post a request to yourself."
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/friends/requests [post]
func (ah *userManagementHandler) PostFriendRequestByEmail(ctx *gin.Context) {
	logrus.Debugf("API-Handler -> GetCurrentUser -> Call to PostFriendRequestByEmail endoint...")

	rawBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, InternalServerErrorMessage)
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cannot read body of request"))
		return
	}

	type postFriendRequestViaEmailRequest struct {
		FriendEmail string `json:"friend_email"`
	}
	var bodyData postFriendRequestViaEmailRequest
	err = json.Unmarshal(rawBody, &bodyData)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Form data is incorrect/invalid.")
		_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("bad request missing friend email"))
		return
	}

	tokenClaims, ok := auth.GetClaims(ctx)
	if !ok {
		ctx.String(http.StatusUnauthorized, "unauthorized")
		_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
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
