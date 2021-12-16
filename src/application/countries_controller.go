package application

import (
	"net/http"
	"world-cup-api/src/domain"
	"world-cup-api/src/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var country domain.Country

	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	countryRepository := repository.GetContryRespositoryInstance()

	countryRepository.Create(&country)

	c.JSON(http.StatusOK, country)
}

func List(c *gin.Context) {
	var countries []domain.Country

	countryRepository := repository.GetContryRespositoryInstance()

	countries = countryRepository.List()

	c.JSON(http.StatusOK, countries)
}
