package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sigma-test/internal/app"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingPongIntegration(t *testing.T) {
	t.Run("should return ping on pong", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			require.NoError(t, err)
		}

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, `{"message":"pong"}`, res.Body.String())
	})
}

func TestSignupIntegration(t *testing.T) {
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()

		requestBody := []byte(`{"email": "test", "password": "test", "role": "user"}`)
		body := bytes.NewReader(requestBody)
		req, err := http.NewRequest("POST", "/api/signup", body)
		if err != nil {
			require.NoError(t, err)
		}

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/api/signup", nil)
		if err != nil {
			require.NoError(t, err)
		}

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 if user exists", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()

		jsonBody := []byte(`{"email": "test@test.com", "password": "test123", "role": "user"}`)
		bodyReader := bytes.NewReader(jsonBody)
		req, err := http.NewRequest("POST", "/api/signup", bodyReader)
		if err != nil {
			require.NoError(t, err)
		}

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
}
