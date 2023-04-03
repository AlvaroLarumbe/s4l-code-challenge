package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"s4l-code-challenge/domain/model"
	"testing"
	"time"
)

type MaximizeCalculateUseCaseTestSuite struct {
	suite.Suite

	useCase MaximizeCalculateUseCase
}

func TestMaximizeCalculateUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(MaximizeCalculateUseCaseTestSuite))
}

func (suite *MaximizeCalculateUseCaseTestSuite) SetupTest() {
	suite.useCase = NewMaximizeCalculateUseCase(NewStatsCalculateUseCase())
}

func (suite *MaximizeCalculateUseCaseTestSuite) TestExecuteWhenSomeRequestsAreGiven() {
	actual := suite.useCase.Execute([]model.BookingPetition{
		{
			PetitionId:  "::request-id-1::", // Profit per night = 8.4
			CheckIn:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			Nights:      5,
			SellingRate: 200,
			Margin:      21,
		},
		{
			PetitionId:  "::request-id-2::", // Profit per night = 1.95
			CheckIn:     time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
			Nights:      4,
			SellingRate: 156,
			Margin:      5,
		},
		{
			PetitionId:  "::request-id-3::", // Profit per night = 2.25
			CheckIn:     time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
			Nights:      4,
			SellingRate: 150,
			Margin:      6,
		},
		{
			PetitionId:  "::request-id-4::", // Profit per night = 12.525
			CheckIn:     time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
			Nights:      4,
			SellingRate: 167,
			Margin:      30,
		},
	})

	expected := model.Maximized{
		PetitionIds: []string{
			"::request-id-1::",
			"::request-id-4::",
		},
		TotalProfit: 92.1,
		AvgNight:    10.4625,
		MinNight:    8.4,
		MaxNight:    12.525,
	}

	assert.Equal(suite.T(), expected, actual)
}
