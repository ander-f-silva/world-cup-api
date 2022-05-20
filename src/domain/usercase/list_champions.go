package usercase

import (
	"world-cup-api/src/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ListChampions interface {
	FindAll() []entity.Champion
}

type ListingChampions struct {
}

func ListChampionsFactory() ListChampions {
	return ListingChampions{}
}

func createConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Champion{})

	return db
}

func (listingChampion ListingChampions) FindAll() []entity.Champion {
	var champions []entity.Champion

	sqlite := createConnection()

	db, _ := sqlite.DB()

	sqlite.Find(&champions)

	db.Close()

	return champions
}
