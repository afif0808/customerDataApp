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
	testRecorder := httptest.NewRecorder()
	testRouter := mux.NewRouter()

	testRouter.Handle("/testWithoutTemplate", ServeViewController("../views/pelanggan.html", ""))
	testRouter.Handle("/testWithTemplate", ServeViewController("../views/pelanggan.html", "../views/template.html"))

	//serve view without template
	testRequestWithoutTemplate := httptest.NewRequest("GET", "http://localhost:123/testWithoutTemplate", nil)
	testRouter.ServeHTTP(testRecorder, testRequestWithoutTemplate)
	assert.Equal(t, successExpectation, testRecorder.Code)

	//serve view with template
	testRequestWithTemplate := httptest.NewRequest("GET", "http://localhost:123/testWithTemplate", nil)
	testRouter.ServeHTTP(testRecorder, testRequestWithTemplate)
	assert.Equal(t, successExpectation, testRecorder.Code)

}
