package repository

import (
	"world-cup-api/src/domain/entity"

	"world-cup-api/src/infrastructure/dao"
	"world-cup-api/src/infrastructure/sqllite"
)

type ChampionRepository interface {
	Create(*entity.Champion)
	List() []entity.Champion
	FindByYear(year int) *entity.Champion
}

func GetChampionRespositoryInstance() ChampionRepository {
	return dao.ChampionDAO{
		DbSqlLite: sqllite.GetInstance(),
	}
}
