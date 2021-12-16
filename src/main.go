package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Country struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
}

func createConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Country{})

	return db
}

func main() {
	router := gin.Default()

	router.POST("/countries", func(c *gin.Context) {

		var country Country

		if err := c.ShouldBindJSON(&country); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sqlLite := createConnection()

		db, _ := sqlLite.DB()

		sqlLite.Create(&country)

		db.Close()

		c.JSON(http.StatusOK, country)
	})

	router.GET("/countries", func(c *gin.Context) {
		sqlLite := createConnection()

		var countries []Country

		// countries[0] = Country{
		// 	Name: "Brasil",
		// }

		// countries[1] = Country{
		// 	Name: "Argentina",
		// }

		db, _ := sqlLite.DB()

		sqlLite.Find(&countries)

		db.Close()

		c.JSON(http.StatusOK, countries)
	})

	router.Run(":3000")
}
