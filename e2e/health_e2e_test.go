package e2e

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	server := createServer()
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	object := e.GET("/health").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	object.Keys().ContainsOnly("status", "version")
	object.Value("status").IsEqual("ok")
	object.Value("version").IsEqual(1)
}
