package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/card"
	"ygodraft/backend/config"
	"ygodraft/backend/ygo"
)

func main() {
	err := startProgramm()
	if err != nil {
		logrus.Error(err)
	}
}

func startProgramm() error {
	client, err := ygo.NewYgoClientWithCache()
	if err != nil {
		return fmt.Errorf("failed to create new ygoCLientCache: %w", err)
	}
	defer client.Close()

	ygoContext, err := config.NewYgoContext("config.yaml", client)
	if err != nil {
		return fmt.Errorf("failed to read config config.yaml: %w", err)
	}

	err = configureLogger(ygoContext)
	if err != nil {
		return fmt.Errorf("failed to configure logger: %w", err)
	}

	err = synchCards(client, ygoContext)

	router, err := setupRouter(ygoContext)
	if err != nil {
		return fmt.Errorf("failed to setup router: %w", err)
	}

	err = router.Run(fmt.Sprintf(":%d", ygoContext.Port))
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func synchCards(ygoClient *ygo.YgoClientWithCache, ygoContext *config.YGOContext) error {
	if ygoContext.SyncAtStartup {
		logrus.Info("Startup -> Detected data sync at startup")

		err := ygoClient.Sync()
		if err != nil {
			return fmt.Errorf("failed to sync cards: %w", err)
		}

		logrus.Info("Startup -> Sync finished")
	}

	return nil
}

func setupRouter(ygoContext *config.YGOContext) (*gin.Engine, error) {
	router := gin.Default()
	router.BasePath()

	router.LoadHTMLFiles("ui/build/index.html")
	publicAPI := router.Group(ygoContext.ContextPath)
	publicAPI.GET("/", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		//c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	publicAPI.StaticFile("favicon.ico", "ui/build/favicon.ico")
	publicAPI.Static("static", "ui/build/static")

	apiV1Group := router.Group("api/v1")
	err := card.SetupAPI(apiV1Group, ygoContext)
	if err != nil {
		return nil, fmt.Errorf("failed to setup api: %w", err)
	}

	return router, nil
}

func configureLogger(ygoContext *config.YGOContext) error {
	logLevel, err := logrus.ParseLevel(ygoContext.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to parse log level %s: %w", ygoContext.LogLevel, err)
	}

	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	return nil
}
