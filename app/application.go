package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
var(
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	fmt.Println("App Started !!")

	router.Run(":8080")


}