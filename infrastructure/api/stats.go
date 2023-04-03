package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"math"
	"net/http"
	"s4l-code-challenge/domain/model"
)

type statsResponse struct {
	AvgNight float64 `json:"avg_night"`
	MinNight float64 `json:"min_night"`
	MaxNight float64 `json:"max_night"`
}

func (s Server) statsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		s.writeNotFoundError(response)

		return
	}

	requestData, err := io.ReadAll(request.Body)

	if err != nil {
		log.Error().Err(err).Msg("Cannot read request body")
	}

	var requestList []BookingPetitionRequest
	err = json.Unmarshal(requestData, &requestList)

	if err != nil {
		log.Error().Err(err).Msg("Cannot unmarshal request body")
	}

	err = validateRequest(s, response, requestList)

	if err != nil {
		// API error already handled in validator function
		return
	}

	var petitionList []model.BookingPetition

	for _, petitionRequest := range requestList {
		petitionModel, err := petitionRequest.toModel()

		if err != nil {
			log.Error().Err(err).Msg("Cannot map BookingPetitionRequest")
			s.writeInternalServerError(response)

			return
		}

		petitionList = append(petitionList, *petitionModel)
	}

	stats := s.statsCalculateUseCase.Execute(petitionList)

	s.writeData(
		response,
		http.StatusOK,
		statsResponse{
			AvgNight: math.Floor(stats.AvgNight*100) / 100,
			MinNight: math.Floor(stats.MinNight*100) / 100,
			MaxNight: math.Floor(stats.MaxNight*100) / 100,
		},
	)
}
