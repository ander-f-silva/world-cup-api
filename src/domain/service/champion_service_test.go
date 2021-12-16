package service

import (
	"fmt"

	"testing"
	"world-cup-api/src/domain/entity"
	"world-cup-api/src/infrastructure/dao"

	"github.com/stretchr/testify/assert"
)

func TestChampionBusinessCreate(t *testing.T) {
	type testCase struct {
		title              string
		paramDummy         *entity.Champion
		answerFindYearMock *entity.Champion
		expected           error
	}

	testCases := []testCase{
		{
			title:              "Should create the champion wiht success",
			paramDummy:         &entity.Champion{Country: "Brasil", Year: 1994},
			answerFindYearMock: &entity.Champion{ID: 0},
			expected:           nil,
		},
		{
			title:              "Should not create the champion already registed",
			paramDummy:         &entity.Champion{Country: "Fraça", Year: 1998},
			answerFindYearMock: &entity.Champion{ID: 2},
			expected:           fmt.Errorf("champion %d already register", 2),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			championMockDAO := new(dao.ChampionMockDAO)

			championMockDAO.On("Create", tc.paramDummy.Year).Return(tc.answerFindYearMock)

			championBusiness := ChampionBusiness{
				championRepository: championMockDAO,
			}

			result := championBusiness.Create(tc.paramDummy)

			assert := assert.New(t)

			if assert.NotNil(result) {
				assert.Equal(tc.expected, result)
			} else {
				assert.Nil(result)
			}
		})
	}
}
