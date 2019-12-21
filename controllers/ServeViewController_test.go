package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestServeViewController(t *testing.T) {
	successExpectation := http.StatusOK
	testRequest := httptest.NewRequest("GET", "http://localhost:8080/test", nil)
	testRecorder := httptest.NewRecorder()
	testRouter := mux.NewRouter()
	testRouter.Handle("/test", ServeViewController("/home.html", ""))
	testRouter.ServeHTTP(testRecorder, testRequest)
	assert.Equal(t, successExpectation, testRecorder.Code)
}
