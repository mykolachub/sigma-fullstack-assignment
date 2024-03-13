package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sigma-test/internal/app"
	"sigma-test/internal/controller"
	"sigma-test/internal/entity"
	"sigma-test/internal/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignupWithMocks(t *testing.T) {
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		r := gin.New()

		userService := &MockUserService{MockDB: nil}
		controller.InitUserHandler(r, userService)

		requestBody := `{"email": "test@example.com", "password": "password", "role": "user"}`
		body := bytes.NewBufferString(requestBody)

		req, err := http.NewRequest("POST", "/api/signup", body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r := gin.New()

		userService := &MockUserService{MockDB: nil}
		controller.InitUserHandler(r, userService)

		req, err := http.NewRequest("POST", "/api/signup", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 if user exists", func(t *testing.T) {
		r := gin.New()

		// Emutaling real users in database
		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test123", Role: "user"}
		userService := &MockUserService{MockDB: []entity.User{mockUser}}
		controller.InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(mockUser)
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("POST", "/api/signup", body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
}

// Testing signup handler with SetupRouter
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

func TestLoginWithMocks(t *testing.T) {
	t.Run("should return 200 on successful login", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		userService := &MockUserService{MockDB: []entity.User{mockUser}}
		controller.InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(mockUser)
		body := bytes.NewBuffer(requestBody)
		req, err := http.NewRequest("POST", "/api/login", body)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r := gin.New()

		userService := &MockUserService{MockDB: nil}
		controller.InitUserHandler(r, userService)

		req, err := http.NewRequest("POST", "/api/login", nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 on invalid credantials", func(t *testing.T) {
		r := gin.New()

		// No users in mocked database, login impossible
		mockUser := entity.User{}
		userService := &MockUserService{MockDB: nil}
		controller.InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(mockUser)
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("POST", "/api/login", body)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
}

func TestMeWithMocks(t *testing.T) {
	t.Run("should return 200 on success", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)

		userService := &MockUserService{MockDB: []entity.User{mockUser}}
		controller.InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/me", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		t.Log(res.Body.String())
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken("INVALID_ID", "INVALID_ROLE")

		userService := &MockUserService{MockDB: []entity.User{mockUser}}
		controller.InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/me", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		t.Log(res.Body.String())
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}

		userService := &MockUserService{MockDB: []entity.User{mockUser}}
		controller.InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/me", nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		t.Log(res.Body.String())
		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

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
