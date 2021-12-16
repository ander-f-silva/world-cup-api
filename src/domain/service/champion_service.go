package service

import (
	"fmt"
	"log"
	"world-cup-api/src/domain/entity"
	"world-cup-api/src/domain/repository"
)

type ChampionService interface {
	Create(*entity.Champion) error
	List() []entity.Champion
}

type ChampionBusiness struct {
	championRepository repository.ChampionRepository
}

func GetChampionServiceInstance() ChampionService {
	return ChampionBusiness{
		championRepository: repository.GetChampionRespositoryInstance(),
	}
}

func (championBusiness ChampionBusiness) Create(champion *entity.Champion) error {
	championFound := championBusiness.championRepository.FindByYear(champion.Year)

	log.Printf("[event: create] Champions found: %v", championFound)

	if championFound.IsExist() {
		return fmt.Errorf("champion %d already register ", champion.Year)
	}

	championBusiness.championRepository.Create(champion)

	return nil
}

func (championBusiness ChampionBusiness) List() []entity.Champion {
	var champions []entity.Champion

	champions = championBusiness.championRepository.List()

	return champions
}
