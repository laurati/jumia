package tests

import (
	"jumia/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPingRoute(t *testing.T) {
	router := routes.LoadRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// /countries -> Error: unable to open database file: no such file or directory (500)
