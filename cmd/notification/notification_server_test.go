package main

import (
	"os"
	"testing"

	"net/http"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

var notificationServer *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)

	var err error
	notificationServer, err = initializeNotificationServer()
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
		return
	}

	os.Exit(m.Run())
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	notificationServer.ServeHTTP(rr, req)

	return rr
}

func TestHealthCheck(t *testing.T) {
	if testing.Short() {
		t.Skip("skip tests")
		return
	}

	tt := []struct {
		name               string
		expectedStatusCode int
		expectedResult     string
	}{
		{"regular", http.StatusOK, "OK"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/internal/health", nil)
			if err != nil {
				t.Error(err)
			}

			response := executeRequest(req)

			assert.Equal(t, response.Code, tc.expectedStatusCode)
			assert.Equal(t, response.Body.String(), tc.expectedResult)
		})
	}
}

func TestVersionCheck(t *testing.T) {
	if testing.Short() {
		t.Skip("skip tests")
		return
	}

	tt := []struct {
		name               string
		expectedStatusCode int
		expectedResult     string
	}{
		{"regular", http.StatusOK, "0.0.0"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/internal/version", nil)
			if err != nil {
				t.Error(err)
			}

			response := executeRequest(req)

			assert.Equal(t, response.Code, tc.expectedStatusCode)
			assert.Equal(t, response.Body.String(), tc.expectedResult)
		})
	}
}
