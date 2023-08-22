package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/model"
)

func SetupAPI(router *gin.RouterGroup, authClient model.YGOJwtAuthClient, client *ygo.YgoClientWithCache, usermgtClient model.UsermgtClient) error {
	// unprotected access for login and stuff
	SetupUnprotectedAPI(router, authClient, client, usermgtClient)

	// access for authenticated default user
	userGroup := router.Use(auth.PermissionMiddleware(authClient, false))
	SetupAuthenticatedUserApi(userGroup, client)

	// access for authenticated admin user
	adminGroup := router.Use(auth.PermissionMiddleware(authClient, true))
	SetupAuthenticatedAdminApi(adminGroup, client)

	return nil
}

func SetupUnprotectedAPI(router gin.IRoutes, authClient model.YGOJwtAuthClient, client *ygo.YgoClientWithCache, usermgtClient model.UsermgtClient) {
	authHandler := newAuthenticationHandler(authClient, usermgtClient)
	router.POST("login", authHandler.Login)

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

func SetupAuthenticatedUserApi(router gin.IRoutes, _ *ygo.YgoClientWithCache) {
	router.GET("usertest", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "user test")
	})
}

func SetupAuthenticatedAdminApi(router gin.IRoutes, _ *ygo.YgoClientWithCache) {
	router.GET("admintest", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "admin test")
	})
}
