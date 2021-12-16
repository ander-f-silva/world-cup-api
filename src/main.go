package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Champion struct {
	ID      uint   `json:"id,omitempty" gorm:"primaryKey"`
	Country string `json:"country" binding:"required"`
	Year    int    `json:"year" binding:"required"`
}

func createConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Champion{})

	return db
}

func main() {
	router := gin.Default()

	router.POST("/champions", func(c *gin.Context) {

		var champion Champion

		if err := c.ShouldBindJSON(&champion); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sqlLite := createConnection()

		db, _ := sqlLite.DB()

		sqlLite.Create(&champion)

		db.Close()

		c.JSON(http.StatusOK, champion)
	})

	router.GET("/champions", func(c *gin.Context) {
		sqlLite := createConnection()

		var champions []Champion

		// champions[0] = Champion{
		// 	Country: "Brasil",
		// 	Year: 1970,
		// }

		// champions[1] = Champion{
		// 	Country: "Argentina",
		// 	Year: 1986,
		// }

		db, _ := sqlLite.DB()

		sqlLite.Find(&champions)

		db.Close()

		c.JSON(http.StatusOK, champions)
	})

	router.Run(":3000")
}
