package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"s4l-code-challenge/domain/model"
	"testing"
	"time"
)

type StatsCalculateUseCaseTestSuite struct {
	suite.Suite

	useCase StatsCalculateUseCase
}

func TestStatsCalculateUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(StatsCalculateUseCaseTestSuite))
}

func (suite *StatsCalculateUseCaseTestSuite) SetupTest() {
	suite.useCase = NewStatsCalculateUseCase()
}

func (suite *StatsCalculateUseCaseTestSuite) TestExecuteWhenEmptyListIsGiven() {
	actual := suite.useCase.Execute([]model.BookingPetition{})

	expected := model.Stats{}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *StatsCalculateUseCaseTestSuite) TestExecuteWhenSomeRequestsAreGiven() {
	actual := suite.useCase.Execute([]model.BookingPetition{
		{
			PetitionId:  "::request-id-1::", // Profit per night = 13.44
			CheckIn:     time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			Nights:      3,
			SellingRate: 224,
			Margin:      18,
		},
		{
			PetitionId:  "::request-id-2::", // Profit per night = 26
			CheckIn:     time.Date(2023, 4, 3, 0, 0, 0, 0, time.UTC),
			Nights:      2,
			SellingRate: 260,
			Margin:      20,
		},
		{
			PetitionId:  "::request-id-3::", // Profit per night = 8.184
			CheckIn:     time.Date(2023, 3, 30, 0, 0, 0, 0, time.UTC),
			Nights:      5,
			SellingRate: 186,
			Margin:      22,
		},
	})

	expected := model.Stats{
		AvgNight: 15.874666666666664,
		MinNight: 8.184000000000001,
		MaxNight: 26,
	}

	assert.Equal(suite.T(), expected, actual)
}
