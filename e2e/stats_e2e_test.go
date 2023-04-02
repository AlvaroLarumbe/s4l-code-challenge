package e2e

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func TestStatsEndpoint(t *testing.T) {
	server := createServer()
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	request := []map[string]interface{}{
		{
			"request_id":   "bookata_XY123",
			"check_in":     "2020-01-01",
			"nights":       5,
			"selling_rate": 200,
			"margin":       20,
		},
		{
			"request_id":   "kayete_PP234",
			"check_in":     "2020-01-04",
			"nights":       4,
			"selling_rate": 156,
			"margin":       22,
		},
	}

	object := e.POST("/stats").
		WithJSON(request).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	object.Keys().ContainsOnly("avg_night", "min_night", "max_night")
	object.Value("avg_night").IsEqual(8.28)
	object.Value("min_night").IsEqual(8)
	object.Value("max_night").IsEqual(8.58)
}
