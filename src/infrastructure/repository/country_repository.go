package repository

import (
	"world-cup-api/src/domain"
	"world-cup-api/src/infrastructure/sqllite"

	"gorm.io/gorm"
)

type CountryRepository interface {
	Create(domain.Country)
	List() []domain.Country
}

type CountryDAO struct {
	dbSqlLite *gorm.DB
}

func GetContryRespositoryInstance() CountryDAO {
	return CountryDAO{
		dbSqlLite: sqllite.GetInstance(),
	}
}

func (dao CountryDAO) Create(country *domain.Country) {
	sqlLite := sqllite.GetInstance()

	db, _ := sqlLite.DB()

	sqlLite.Create(country)

	db.Close()
}

func (dao CountryDAO) List() []domain.Country {
	var countries []domain.Country

	sqlLite := sqllite.GetInstance()

	db, _ := sqlLite.DB()

	sqlLite.Find(&countries)

	db.Close()

	return countries
}
