package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/model"
)

const InternalServerErrorMessage = "Internal server error. Check server logs for more information."

// SetupAPI creates the routes for the api endpoints.
func SetupAPI(router *gin.RouterGroup, authClient model.YGOJwtAuthClient, client *ygo.YgoClientWithCache, usermgtClient model.UsermgtClient) error {
	SwaggerInfo.BasePath = "/api/v1"

	authHandler := newAuthenticationHandler(authClient, usermgtClient)
	usermgtHandler := newUserManagementHandler(usermgtClient)

	// unprotected access for login and stuff
	setupUnprotectedAPI(router, client, authHandler)

	// access for authenticated default user
	userGroup := router.Use(auth.PermissionMiddleware(authClient, false))
	setupAuthenticatedUserApi(userGroup, client, usermgtHandler)

	// access for authenticated admin user
	adminGroup := router.Use(auth.PermissionMiddleware(authClient, true))
	setupAuthenticatedAdminApi(adminGroup, client, usermgtHandler)

	return nil
}

func setupUnprotectedAPI(router gin.IRoutes, client *ygo.YgoClientWithCache, handler *authenticationHandler) {
	router.POST("login", handler.Login)

	cardRetriever := newYgoRetrieveHandler(client)

	router.GET("cards", cardRetriever.GetCards)
	router.GET("cards/:id", cardRetriever.GetCard)
	router.GET("cards/random", cardRetriever.GetRandomCards)
	router.GET("sets", cardRetriever.GetSets)
	router.GET("sets/:code", cardRetriever.GetSet)
	router.GET("sets/:code/cards", cardRetriever.GetSetCards)
}

func setupAuthenticatedUserApi(router gin.IRoutes, _ *ygo.YgoClientWithCache, usermgtHandler *userManagementHandler) {
	router.GET("user", usermgtHandler.GetCurrentUser)
	router.GET("user/friends", usermgtHandler.GetFriends)
	router.GET("user/friends/requests", usermgtHandler.GetFriendRequests)
	router.POST("user/friends/requests/:id", usermgtHandler.PostFriendRequest)
	router.POST("user/friends/requests", usermgtHandler.PostFriendRequestByEmail)
}

func setupAuthenticatedAdminApi(router gin.IRoutes, _ *ygo.YgoClientWithCache, usermgtHandler *userManagementHandler) {
	router.GET("users", usermgtHandler.GetUsers)
	router.POST("users", usermgtHandler.PostUsers)
	router.DELETE("users", usermgtHandler.DeleteUser)
}

// GetQueryParameterInt retrieves an integer query parameter from the context.
func GetQueryParameterInt(ctx *gin.Context, parameterName string) (int, error) {
	raw, exists := ctx.GetQuery(parameterName)
	if !exists {
		return 0, fmt.Errorf("required query parameter [%s] does not exist", parameterName)
	}

	parameterValue, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("required query parameter [%s=%s] does not seem to be a valid integer", parameterName, raw)
	}

	return parameterValue, nil
}
