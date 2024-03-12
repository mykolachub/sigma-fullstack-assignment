package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"sigma-test/internal/app"
	"sigma-test/internal/controller"
	"sigma-test/internal/entity"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockUserService struct {
	MockDB []entity.User
}

func (m *MockUserService) DeleteUser(email string) error {
	return nil
}

func (m *MockUserService) GetAllUsers() ([]response.User, error) {
	return []response.User{}, nil
}

func (m *MockUserService) GetUserById(id string) (response.User, error) {
	return response.User{}, nil
}

func (m *MockUserService) Login(body request.User) (string, error) {
	return "", nil
}

func (m *MockUserService) UpdateUser(id string, user request.User) (response.User, error) {
	return response.User{}, nil
}

func (m *MockUserService) GetUserByEmail(email string) (response.User, error) {
	return response.User{}, nil
}

func (m *MockUserService) CreateUser(user request.User) (response.User, error) {
	return response.User{}, nil
}

// SignUp implements controller.UserService.
func (m *MockUserService) SignUp(body request.User) (response.User, error) {
	for _, v := range m.MockDB {
		log.Println(v.Email)
		log.Println(body.Email)

		if v.Email == body.Email {
			return response.User{}, errors.New("user already exists")
		}
	}
	return response.User{ID: "test", Email: "test@test.com", Password: "password", Role: "user"}, nil
}

// Testing signup handler with mocking user service
func TestSignupWithMocks(t *testing.T) {
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		r := gin.New()

		userService := &MockUserService{MockDB: nil}

		controller.InitUserHandler(r, userService)

		requestBody := `{"email": "test@example.com", "password": "password", "role": "user"}`
		body := bytes.NewBufferString(requestBody)

		req, err := http.NewRequest("POST", "/api/signup", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %v", http.StatusOK, res.Code)
		}
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r := gin.New()

		userService := &MockUserService{MockDB: nil}

		controller.InitUserHandler(r, userService)

		req, err := http.NewRequest("POST", "/api/signup", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d but got %d", http.StatusBadRequest, res.Code)
		}
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
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		if res.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status code %d but got %d", http.StatusUnprocessableEntity, res.Code)
		}
	})
}

// Testing signup handler with SetupRouter
func TestSignup(t *testing.T) {
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()

		requestBody := []byte(`{"email": "test", "password": "test", "role": "user"}`)
		body := bytes.NewReader(requestBody)
		req, err := http.NewRequest("POST", "/api/signup", body)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, res.Code)
		}
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/api/signup", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d but got %d", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("should return 422 if user exists", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()

		jsonBody := []byte(`{"email": "test@test.com", "password": "test123", "role": "user"}`)
		bodyReader := bytes.NewReader(jsonBody)
		req, err := http.NewRequest("POST", "/api/signup", bodyReader)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(res, req)

		// Check the response status code
		if res.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status code %d but got %d", http.StatusUnprocessableEntity, res.Code)
		}
	})
}

func TestPingPong(t *testing.T) {
	t.Run("should return ping on pong", func(t *testing.T) {
		router := app.SetupRouter()
		res := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, res.Code)
		}

		expectedBody := `{"message":"pong"}`
		if res.Body.String() != expectedBody {
			t.Errorf("expected body %s but got %s", expectedBody, res.Body.String())
		}
	})

}
