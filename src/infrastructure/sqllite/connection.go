package sqllite

import (
	"world-cup-api/src/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&domain.Country{})

	return db
}
