package main

import (
	championApi "world-cup-api/src/application/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/champions", championApi.Create)
	router.GET("/champions", championApi.List)

	router.Run(":3000")
}
