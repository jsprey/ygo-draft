package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"ygodraft/backend/api"
	"ygodraft/backend/client/auth"
	"ygodraft/backend/client/postgresql"
	"ygodraft/backend/client/usermgt"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
	"ygodraft/backend/setup"
	"ygodraft/backend/synch"
)

func main() {
	err := startProgram()
	if err != nil {
		logrus.Error(err)
	}
}

func startProgram() error {
	ygoCtx, err := config.NewYgoContext("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to read config config.yaml: %w", err)
	}

	err = configureLogger(ygoCtx)
	if err != nil {
		return fmt.Errorf("failed to configure logger: %w", err)
	}

	logrus.Debugf("Startup -> Current Config [%+v]", ygoCtx)

	dbClient, err := setupDB(ygoCtx)
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}
	defer dbClient.PoolConnection.Close()

	usermgtClient, err := usermgt.NewUsermgtClient(dbClient)
	if err != nil {
		return err
	}

	ygoClient, err := setupYgoClient(ygoCtx, dbClient)
	if err != nil {
		return fmt.Errorf("failed to setup ygo client: %w", err)
	}

	router, err := setupRouter(ygoCtx, ygoClient, usermgtClient)
	if err != nil {
		return fmt.Errorf("failed to setup router: %w", err)
	}

	err = router.Run(fmt.Sprintf(":%d", ygoCtx.Port))
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func setupDB(ygoCtx *config.YgoContext) (*postgresql.PostgresClient, error) {
	logrus.Info("Startup -> Setup database")
	client, err := postgresql.NewPostgresClient(ygoCtx.DatabaseContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create new db client: %w", err)
	}

	usermgtClient, err := usermgt.NewUsermgtClient(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create new usermgt client: %w", err)
	}

	databaseSetup := setup.NewDatabaseSetup(client, usermgtClient)
	err = databaseSetup.Setup()
	if err != nil {
		return nil, fmt.Errorf("failed to perform database setup: %w", err)
	}

	return client, nil
}

func setupYgoClient(ygoCtx *config.YgoContext, dbClient model.DatabaseClient) (*ygo.YgoClientWithCache, error) {
	client, err := ygo.NewYgoClientWithCache(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ygoCLientCache: %w", err)
	}

	if ygoCtx.SyncAtStartup {
		logrus.Info("Startup -> Detected data sync at startup")

		dataSyncher, err := synch.NewYgoDataSyncher(client, ygoCtx)
		if err != nil {
			return nil, err
		}

		err = dataSyncher.Sync()
		if err != nil {
			return nil, fmt.Errorf("failed to sync cards: %w", err)
		}

		logrus.Info("Startup -> Sync finished")
	}

	return client, nil
}

func setupRouter(ygoCtx *config.YgoContext, client *ygo.YgoClientWithCache, usermgtClient model.UsermgtClient) (*gin.Engine, error) {
	authClient := auth.NewYgoJwtAuthClient(ygoCtx.AuthenticationContext.JWTSecretKey)

	router := gin.Default()
	router.BasePath()
	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE")
		context.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204) // Respond with a No Content status for preflight requests
			return
		}

		context.Next()
	})

	router.LoadHTMLFiles("build/ui/index.html")
	publicAPI := router.Group(ygoCtx.ContextPath)
	publicAPI.GET("/", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	publicAPI.StaticFile("favicon.ico", "build/ui/favicon.ico")
	publicAPI.Static("static", "build/ui/static")
	publicAPI.Static("/images/cards", "./imageStore")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiV1Group := publicAPI.Group("api/v1")
	err := api.SetupAPI(apiV1Group, authClient, client, usermgtClient)
	if err != nil {
		return nil, fmt.Errorf("failed to setup api: %w", err)
	}

	return router, nil
}

func configureLogger(ygoCtx *config.YgoContext) error {
	logLevel, err := logrus.ParseLevel(ygoCtx.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to parse log level %s: %w", ygoCtx.LogLevel, err)
	}

	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	return nil
}
