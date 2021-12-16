package main

import (
	"world-cup-api/src/application"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/countries", application.Create)

	router.GET("/countries", application.List)

	router.Run(":3000")
}
