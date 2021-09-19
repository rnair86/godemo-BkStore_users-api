package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rnair86/godemo-BkStore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Starting Application")
	err := router.Run(":8080")
	if err != nil {
		logger.Fatal("unable to start application ", err)
		return
	}
}
