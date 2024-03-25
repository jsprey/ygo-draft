package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/model"
)

const InternalServerErrorMessage = "Internal server error. Check server logs for more information."

func SetupAPI(router *gin.RouterGroup, authClient model.YGOJwtAuthClient, client *ygo.YgoClientWithCache, usermgtClient model.UsermgtClient) error {
	authHandler := newAuthenticationHandler(authClient, usermgtClient)

	// unprotected access for login and stuff
	SetupUnprotectedAPI(router, client, authHandler)

	// access for authenticated default user
	userGroup := router.Use(auth.PermissionMiddleware(authClient, false))
	SetupAuthenticatedUserApi(userGroup, client, authHandler)

	// access for authenticated admin user
	adminGroup := router.Use(auth.PermissionMiddleware(authClient, true))
	SetupAuthenticatedAdminApi(adminGroup, client, authHandler)

	return nil
}

func SetupUnprotectedAPI(router gin.IRoutes, client *ygo.YgoClientWithCache, handler *authenticationHandler) {
	router.POST("login", handler.Login)

	cardRetriever := CardRetrieveHandler{
		YGOClient: client,
	}

	router.GET("cards", cardRetriever.GetCards)
	router.GET("cards/:id", cardRetriever.GetCard)
	router.GET("randomCard", cardRetriever.GetRandomCard)
	router.GET("randomCards", cardRetriever.GetRandomCards)
	router.GET("sets", cardRetriever.GetSets)
	router.GET("sets/:code", cardRetriever.GetSet)
	router.GET("sets/:code/cards", cardRetriever.GetSetCards)
}

func SetupAuthenticatedUserApi(router gin.IRoutes, _ *ygo.YgoClientWithCache, authHandler *authenticationHandler) {
	router.GET("user", authHandler.CurrentUser)
	router.GET("user/friends", authHandler.GetFriends)
	router.GET("user/friends/requests", authHandler.GetFriendRequests)
	router.POST("user/friends/requests/:id", authHandler.PostRequestFriend)
	router.POST("user/friends/requests", authHandler.PostRequestFriendByEmail)
}

func SetupAuthenticatedAdminApi(router gin.IRoutes, _ *ygo.YgoClientWithCache, authHandler *authenticationHandler) {
	router.GET("users", authHandler.GetUsers)
	router.POST("users", authHandler.RegisterUser)
	router.DELETE("users", authHandler.DeleteUser)
}

func ExtractBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header provided")
	}

	if strings.HasPrefix(authHeader, "Bearer ") {
		token := authHeader[len("Bearer "):]
		return token, nil
	}

	return "", fmt.Errorf("no bearer token provided")
}

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
