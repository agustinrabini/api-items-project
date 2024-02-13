package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/agustinrabini/go-toolkit/goutils/logger"

	"github.com/gin-gonic/gin"
)

// Helper function to process a request and test its response.
func executeRequest(router *gin.Engine, method string, relativeURL string, headers map[string]string, body string) *httptest.ResponseRecorder {
	var request *http.Request
	var err error

	url := "http://localhost:8080" + relativeURL

	logger.Infof("Making request %s %s", method, url)

	if body != "" {
		request, err = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		logger.Error("Error executing http request with body to "+url, err)
		return nil
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}
