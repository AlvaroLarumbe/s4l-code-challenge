package api

import "net/http"

type healthResponse struct {
	Status  string `json:"status"`
	Version uint8  `json:"version"`
}

func (s Server) healthHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		s.writeNotFoundError(response)

		return
	}

	s.writeData(
		response,
		http.StatusOK,
		healthResponse{
			Status:  "ok",
			Version: 1,
		},
	)
}
