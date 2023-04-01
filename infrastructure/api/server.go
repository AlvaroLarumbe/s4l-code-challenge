package api

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"s4l-code-challenge/application/usecase"
)

type Server struct {
	listenAddress         string
	statsCalculateUseCase usecase.StatsCalculateUseCase
}

func NewServer(listenAddress string, statsCalculateUseCase usecase.StatsCalculateUseCase) Server {
	return Server{
		listenAddress:         listenAddress,
		statsCalculateUseCase: statsCalculateUseCase,
	}
}

func (s Server) Start() error {
	handler := http.NewServeMux()

	// TODO: improve route handling using any library to handle both methods and paths
	handler.Handle("/health", http.HandlerFunc(s.healthHandler))
	handler.Handle("/stats", http.HandlerFunc(s.statsHandler))

	return http.ListenAndServe(s.listenAddress, handler)
}

func (s Server) writeNotFoundError(response http.ResponseWriter) {
	s.writeError(response, http.StatusNotFound, "Not Found")
}

func (s Server) writeBadRequestError(response http.ResponseWriter) {
	s.writeError(response, http.StatusBadRequest, "Bad Request, please check your data")
}

func (s Server) writeMissingFieldError(response http.ResponseWriter, field string) {
	s.writeError(response, http.StatusBadRequest, fmt.Sprintf("Field %s is missing", field))
}

func (s Server) writeInternalServerError(response http.ResponseWriter) {
	s.writeError(response, http.StatusInternalServerError, "Internal Server Error")
}

func (s Server) writeError(response http.ResponseWriter, code int, message string) {
	errorResponse := ErrorResponse{
		Message: message,
	}
	s.writeData(response, code, errorResponse)
}

func (s Server) writeData(response http.ResponseWriter, code int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)

	bytes, err := json.Marshal(data)

	if err != nil {
		log.Error().Err(err).Msg("Cannot JSON marshal HTTP response")
		return
	}

	_, err = response.Write(bytes)

	if err != nil {
		log.Error().Err(err).Msg("Cannot write HTTP response")
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
