package usecase

import "s4l-code-challenge/domain/model"

type StatsCalculateUseCase struct {
}

func NewStatsCalculateUseCase() StatsCalculateUseCase {
	return StatsCalculateUseCase{}
}

func (u StatsCalculateUseCase) Execute(bookingPetitionList []model.BookingPetition) model.Stats {
	var stats model.Stats
	var profitsPerNight []float64

	for i, petition := range bookingPetitionList {
		profitPerNight := (float64(petition.SellingRate*petition.Margin) / 100) / float64(petition.Nights)
		profitsPerNight = append(profitsPerNight, profitPerNight)

		// Keep minimum
		if i == 0 || profitPerNight < stats.MinNight {
			stats.MinNight = profitPerNight
		}

		// Keep maximum
		if i == 0 || profitPerNight > stats.MaxNight {
			stats.MaxNight = profitPerNight
		}
	}

	// Calculate average
	for _, profitPerNight := range profitsPerNight {
		stats.AvgNight += profitPerNight
	}

	if len(profitsPerNight) > 0 {
		stats.AvgNight = stats.AvgNight / float64(len(profitsPerNight))
	}

	return stats
}
