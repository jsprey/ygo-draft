package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/client/postgresql"
	"ygodraft/backend/client/ygo"
	"ygodraft/backend/config"
	"ygodraft/backend/setup"
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

	dbClient, err := setupDB(ygoCtx)
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}
	defer dbClient.PoolConnection.Close()

	ygoClient, err := setupYgoClient(ygoCtx, dbClient)
	if err != nil {
		return fmt.Errorf("failed to setup ygo client: %w", err)
	}

	router, err := setupRouter(ygoCtx, ygoClient)
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
	//client, err := postgresql.NewPosgresqlClient(ygoCtx.DatabaseContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create new db client: %w", err)
	}

	databaseSetup := setup.NewDatabaseSetup(client)
	err = databaseSetup.Setup()
	if err != nil {
		return nil, fmt.Errorf("failed to perform database setup: %w", err)
	}

	return client, nil
}

func setupYgoClient(ygoCtx *config.YgoContext, dbClient cache.DatabaseClient) (*ygo.YgoClientWithCache, error) {
	client, err := ygo.NewYgoClientWithCache(dbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ygoCLientCache: %w", err)
	}

	if ygoCtx.SyncAtStartup {
		logrus.Info("Startup -> Detected data sync at startup")

		dataSyncher := setup.NewYgoDataSyncher(client)
		err = dataSyncher.Sync()
		if err != nil {
			return nil, fmt.Errorf("failed to sync cards: %w", err)
		}

		logrus.Info("Startup -> Sync finished")
	}

	return client, nil
}

func setupRouter(ygoCtx *config.YgoContext, client *ygo.YgoClientWithCache) (*gin.Engine, error) {
	router := gin.Default()
	router.BasePath()

	router.LoadHTMLFiles("ui/build/index.html")
	publicAPI := router.Group(ygoCtx.ContextPath)
	publicAPI.GET("/", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		//c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	publicAPI.StaticFile("favicon.ico", "ui/build/favicon.ico")
	publicAPI.Static("static", "ui/build/static")

	//apiV1Group := router.Group("api/v1")
	//err := api.SetupAPI(apiV1Group, client)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to setup api: %w", err)
	//}

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
