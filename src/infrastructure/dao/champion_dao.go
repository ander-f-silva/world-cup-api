package dao

import (
	"world-cup-api/src/domain/entity"
	"world-cup-api/src/infrastructure/sqllite"

	"gorm.io/gorm"
)

type ChampionDAO struct {
	DbSqlLite *gorm.DB
}

func (dao ChampionDAO) Create(champion *entity.Champion) {
	sqlLite := sqllite.GetInstance()

	db, _ := sqlLite.DB()

	sqlLite.Create(champion)

	db.Close()
}

func (dao ChampionDAO) List() []entity.Champion {
	var champions []entity.Champion

	sqlLite := sqllite.GetInstance()

	db, _ := sqlLite.DB()

	sqlLite.Find(champions)

	db.Close()

	return champions
}

func (dao ChampionDAO) FindByYear(year int) *entity.Champion {
	var champion entity.Champion

	sqlLite := sqllite.GetInstance()

	db, _ := sqlLite.DB()

	sqlLite.First(&champion, "year = ?", year)

	db.Close()

	return &champion
}
