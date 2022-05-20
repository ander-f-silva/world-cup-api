package api

import (
	"net/http"
	"world-cup-api/src/domain/entity"
	"world-cup-api/src/domain/usercase"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context) {
	var champion entity.Champion

	if err := context.ShouldBindJSON(&champion); err != nil {
		context.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addChampions := usercase.AddChampionsFactory()

	addChampions.Effect(&champion)

	context.JSONP(http.StatusOK, champion)
}

func List(context *gin.Context) {
	listChampions := usercase.ListChampionsFactory()

	context.JSONP(http.StatusOK, listChampions.FindAll())
}
