package usercase

import (
	"world-cup-api/src/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AddChampions interface {
	Effect(*entity.Champion)
}

type AddingChampion struct {
}

func AddChampionsFactory() AddChampions {
	return AddingChampion{}
}

func createConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Champion{})

	return db
}

func (addingChampion AddingChampion) Effect(champion *entity.Champion) {
	sqlite := createConnection()

	db, _ := sqlite.DB()

	sqlite.Create(&champion)

	db.Close()
}
