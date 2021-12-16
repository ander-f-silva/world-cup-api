package api

import (
	"net/http"
	"world-cup-api/src/domain/entity"
	"world-cup-api/src/domain/service"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var champion entity.Champion

	if err := c.ShouldBindJSON(&champion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	championService := service.GetChampionServiceInstance()

	err := championService.Create(&champion)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, champion)
}

func List(c *gin.Context) {
	var champions []entity.Champion

	championService := service.GetChampionServiceInstance()

	champions = championService.List()

	c.JSON(http.StatusOK, champions)
}
