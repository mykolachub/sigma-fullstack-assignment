package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sigma-test/internal/entity"
	"sigma-test/internal/mock"
	"sigma-test/internal/request"
	"sigma-test/internal/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignup(t *testing.T) {
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		r := gin.New()

		userService := &mock.MockUserService{MockDB: nil}
		InitUserHandler(r, userService)

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

		userService := &mock.MockUserService{MockDB: nil}
		InitUserHandler(r, userService)

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
		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

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

func TestLogin(t *testing.T) {
	t.Run("should return 200 on successful login", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

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

		userService := &mock.MockUserService{MockDB: nil}
		InitUserHandler(r, userService)

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
		userService := &mock.MockUserService{MockDB: nil}
		InitUserHandler(r, userService)

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

func TestMe(t *testing.T) {
	t.Run("should return 200 on success", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/me", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken("INVALID_ID", "INVALID_ROLE")

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/me", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/me", nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("should return 200 on success", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/users", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/users", nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestGetUsersById(t *testing.T) {
	t.Run("should return 200 on success", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/user?id="+mockUser.ID, nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/user?id="+"INVALID_ID", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("GET", "/api/user?id="+mockUser.ID, nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should return 201 on success", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin}}
		InitUserHandler(r, userService)

		mockNewUser := entity.User{ID: "new", Email: "new@test.com", Password: "test", Role: "user"}
		requestBody, _ := json.Marshal(mockNewUser)
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("POST", "/api/users", body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("should return 403 on non-admin", func(t *testing.T) {
		r := gin.New()

		mockNonAdmin := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockNonAdmin.ID, mockNonAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockNonAdmin}}
		InitUserHandler(r, userService)

		mockNewUser := entity.User{ID: "new", Email: "new@test.com", Password: "test", Role: "user"}
		requestBody, _ := json.Marshal(mockNewUser)
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("POST", "/api/users", body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on user already exists", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin}}
		InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(mockAdmin)
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("POST", "/api/users", body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "admin"}

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("POST", "/api/users", nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("should return 200 on update by owner", func(t *testing.T) {
		r := gin.New()

		mockOwner := entity.User{ID: "owner", Email: "owner@test.com", Password: "owner", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockOwner.ID, mockOwner.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockOwner}}
		InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(request.User{Email: "changed@test.com"})
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("PATCH", "/api/users?id="+mockOwner.ID, body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 200 on update by admin", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "admin", Email: "admin@test.com", Password: "admin", Role: "admin"}
		mockUser := entity.User{ID: "user", Email: "user@test.com", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin, mockUser}}
		InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(request.User{Email: "changed@test.com"})
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("PATCH", "/api/users?id="+mockUser.ID, body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 403 on update by non-owner and non-admin", func(t *testing.T) {
		r := gin.New()

		mockUser1 := entity.User{ID: "user1", Email: "user1@test.com", Password: "user1", Role: "user"}
		mockUser2 := entity.User{ID: "user2", Email: "user2@test.com", Password: "user2", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser1.ID, mockUser1.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser1, mockUser2}}
		InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(request.User{Email: "changed@test.com"})
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("PATCH", "/api/users?id="+mockUser2.ID, body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "admin", Email: "admin@test.com", Password: "admin", Role: "admin"}
		mockUser := entity.User{ID: "user", Email: "user@test.com", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin, mockUser}}
		InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(request.User{Email: "changed@test.com"})
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("PATCH", "/api/users?id="+"INVALID_ID", body)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "user", Email: "user@test.com", Password: "user", Role: "user"}

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		requestBody, _ := json.Marshal(request.User{Email: "changed@test.com"})
		body := bytes.NewBuffer(requestBody)

		req, err := http.NewRequest("PATCH", "/api/users?id="+mockUser.ID, body)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("should return 200 on delete by owner", func(t *testing.T) {
		r := gin.New()

		mockOwner := entity.User{ID: "owner", Email: "owner@test.com", Password: "owner", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockOwner.ID, mockOwner.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockOwner}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("DELETE", "/api/users?id="+mockOwner.ID, nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 200 on delete by admin", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "admin", Email: "admin@test.com", Password: "admin", Role: "admin"}
		mockUser := entity.User{ID: "user", Email: "user@test.com", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin, mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("DELETE", "/api/users?id="+mockUser.ID, nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 403 on delete by non-owner and non-admin", func(t *testing.T) {
		r := gin.New()

		mockUser1 := entity.User{ID: "user1", Email: "user1@test.com", Password: "user1", Role: "user"}
		mockUser2 := entity.User{ID: "user2", Email: "user2@test.com", Password: "user2", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser1.ID, mockUser1.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser1, mockUser2}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("DELETE", "/api/users?id="+mockUser2.ID, nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r := gin.New()

		mockAdmin := entity.User{ID: "admin", Email: "admin@test.com", Password: "admin", Role: "admin"}
		mockUser := entity.User{ID: "user", Email: "user@test.com", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		userService := &mock.MockUserService{MockDB: []entity.User{mockAdmin, mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("DELETE", "/api/users?id="+"INVALID_ID", nil)
		if err != nil {
			require.NoError(t, err)
		}
		req.Header.Add("Authorization", "Bearer "+mockToken)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	// 401 on missing token
	t.Run("should return 401 on missing token", func(t *testing.T) {
		r := gin.New()

		mockUser := entity.User{ID: "user", Email: "user@test.com", Password: "user", Role: "user"}

		userService := &mock.MockUserService{MockDB: []entity.User{mockUser}}
		InitUserHandler(r, userService)

		req, err := http.NewRequest("DELETE", "/api/users?id="+mockUser.ID, nil)
		if err != nil {
			require.NoError(t, err)
		}

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}
