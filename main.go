package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/card"
	"ygodraft/backend/config"
)

func main() {
	err := startProgramm()
	if err != nil {
		logrus.Error(err)
	}
}

func startProgramm() error {
	ygoContext, err := config.ReadConfig("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to read config config.yaml: %w", err)
	}

	err = configureLogger(ygoContext)
	if err != nil {
		return fmt.Errorf("failed to configure logger: %w", err)
	}

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
	card.SetupAPI(apiV1Group, ygoContext)

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
