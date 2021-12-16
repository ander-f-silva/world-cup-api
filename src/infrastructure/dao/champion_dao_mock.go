package dao

import (
	"world-cup-api/src/domain/entity"

	"github.com/stretchr/testify/mock"
)

type ChampionMockDAO struct {
	mock.Mock
}

func (mock *ChampionMockDAO) Create(champion *entity.Champion) {
	mock.Called(champion)
}

func (mock *ChampionMockDAO) List() []entity.Champion {
	args := mock.Called()
	return args.Get(0).([]entity.Champion)
}

func (mock *ChampionMockDAO) FindByYear(year int) *entity.Champion {
	args := mock.Called(year)
	return args.Get(0).(*entity.Champion)
}
