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

//func List(context *gin.Context) {
//	sqlite := createConnection()
//
//	var champion []Champion
//
//	db, _ := sqlite.DB()
//
//	sqlite.Find(&champion)
//
//	db.Close()
//
//	context.JSONP(http.StatusOK, champion)
//}
