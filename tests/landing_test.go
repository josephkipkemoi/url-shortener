package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/server"

	"github.com/stretchr/testify/assert"
)

func TestLandingPage(t *testing.T) {
	router := server.ConnectServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/new", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
