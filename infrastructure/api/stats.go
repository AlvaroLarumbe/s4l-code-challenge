package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type statsResponse struct {
	AvgNight float32 `json:"avg_night"`
	MinNight float32 `json:"min_night"`
	MaxNight float32 `json:"max_night"`
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

	// TODO: calculate stats
	s.writeData(
		response,
		http.StatusOK,
		statsResponse{
			AvgNight: 1,
			MinNight: 2,
			MaxNight: 3,
		},
	)
}
