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
			PetitionId:  "::request-id-1::",
			CheckIn:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			Nights:      5,
			SellingRate: 200,
			Margin:      20,
		},
		{
			PetitionId:  "::request-id-2::",
			CheckIn:     time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
			Nights:      4,
			SellingRate: 156,
			Margin:      5,
		},
		{
			PetitionId:  "::request-id-3::",
			CheckIn:     time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
			Nights:      4,
			SellingRate: 150,
			Margin:      6,
		},
		{
			PetitionId:  "::request-id-4::",
			CheckIn:     time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
			Nights:      4,
			SellingRate: 160,
			Margin:      30,
		},
	})

	expected := model.Maximized{
		PetitionIds: []string{
			"::request-id-1::",
			"::request-id-4::",
		},
		TotalProfit: 88,
		AvgNight:    10,
		MinNight:    8,
		MaxNight:    12,
	}

	assert.Equal(suite.T(), expected, actual)
}
