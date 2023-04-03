package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"s4l-code-challenge/domain/model"
)

type maximizeResponse struct {
	PetitionIds []string `json:"request_ids"`
	TotalProfit float64  `json:"total_profit"`
	AvgNight    float64  `json:"avg_night"`
	MinNight    float64  `json:"min_night"`
	MaxNight    float64  `json:"max_night"`
}

func (s Server) maximizeHandler(response http.ResponseWriter, request *http.Request) {
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

	maximized := s.maximizeCalculateUseCase.Execute(petitionList)

	s.writeData(
		response,
		http.StatusOK,
		maximizeResponse{
			PetitionIds: maximized.PetitionIds,
			TotalProfit: maximized.TotalProfit,
			AvgNight:    maximized.AvgNight,
			MinNight:    maximized.MinNight,
			MaxNight:    maximized.MaxNight,
		},
	)
}
