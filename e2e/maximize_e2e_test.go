package e2e

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func TestMaximizeEndpoint(t *testing.T) {
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
		{
			"request_id":   "atropote_AA930",
			"check_in":     "2020-01-04",
			"nights":       4,
			"selling_rate": 150,
			"margin":       6,
		},
		{
			"request_id":   "acme_AAAAA",
			"check_in":     "2020-01-10",
			"nights":       4,
			"selling_rate": 160,
			"margin":       30,
		},
	}

	object := e.POST("/maximize").
		WithJSON(request).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	object.Keys().ContainsOnly("request_ids", "total_profit", "avg_night", "min_night", "max_night")
	object.Value("request_ids").Array().IsEqual([]string{"bookata_XY123", "acme_AAAAA"})
	object.Value("total_profit").IsEqual(88)
	object.Value("avg_night").IsEqual(10)
	object.Value("min_night").IsEqual(8)
	object.Value("max_night").IsEqual(12)
}
