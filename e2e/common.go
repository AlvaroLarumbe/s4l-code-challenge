package e2e

import (
	"net/http/httptest"
	"s4l-code-challenge/application/usecase"
	"s4l-code-challenge/infrastructure/api"
)

// createServer is the single point to init the app for E2E tests
func createServer() *httptest.Server {
	statsCalculateUseCase := usecase.NewStatsCalculateUseCase()
	serveMux := api.NewServer(
		"",
		statsCalculateUseCase,
		usecase.NewMaximizeCalculateUseCase(statsCalculateUseCase),
	)

	return httptest.NewServer(serveMux.CreateServeMux())
}
