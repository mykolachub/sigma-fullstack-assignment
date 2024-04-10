package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sigma-user/config"
	"sigma-user/internal/controller/mocks"
	"sigma-user/internal/entity"
	"sigma-user/internal/request"
	"sigma-user/internal/response"
	"sigma-user/internal/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeUserService(t *testing.T) (*gin.Engine, *mocks.UserService, UserHandlerConfig) {
	r := gin.New()
	userService := mocks.NewUserService(t)
	userHandlerConfig := UserHandlerConfig{JwtSecret: "test_secret"}
	InitUserHandler(r, userService, userHandlerConfig)

	return r, userService, userHandlerConfig
}

func TestSignup(t *testing.T) {
	signUpRequest := func(t *testing.T, body string) *http.Request {
		req, err := http.NewRequest("POST", "/api/user/signup", bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		return req
	}
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		r, usrSvc, _ := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test123", Role: "user"}
		usrSvc.EXPECT().SignUp(mockUser).Return(response.User(mockUser.ToEntity()), config.SvcFailedCreateUser, nil)

		body, _ := json.Marshal(mockUser)
		req := signUpRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		req := signUpRequest(t, "")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 if user exists", func(t *testing.T) {
		r, usrSvc, _ := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test123", Role: "user"}
		usrSvc.EXPECT().SignUp(mockUser).Return(response.User{}, config.SvcUserExists, errors.New("User already exists"))

		body, _ := json.Marshal(mockUser)
		req := signUpRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
}

func TestLogin(t *testing.T) {
	loginRequest := func(t *testing.T, body string) *http.Request {
		req, err := http.NewRequest("POST", "/api/user/login", bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		return req
	}
	t.Run("should return 200 on successful login", func(t *testing.T) {
		r, usrSvc, _ := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		usrSvc.EXPECT().Login(mockUser).Return("token", config.SvcEmptyMsg, nil)

		body, _ := json.Marshal(mockUser)
		req := loginRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		req := loginRequest(t, "")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 on invalid credantials", func(t *testing.T) {
		r, usrSvc, _ := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		usrSvc.EXPECT().Login(mockUser).Return("", config.SvcInvalidCredentials, errors.New("Invalid Credantials"))

		body, _ := json.Marshal(mockUser)
		req := loginRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
}

func TestMe(t *testing.T) {
	meRequest := func(t *testing.T, token string) *http.Request {
		req, err := http.NewRequest("GET", "/api/users/me", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on success", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role, handCfg.JwtSecret)
		usrSvc.EXPECT().GetUserById(mockUser.ID).Return(mockUser.ToResponse(), config.SvcEmptyMsg, nil)

		req := meRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid token", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockToken, _ := util.GenerateJWTToken("INVALID_ID", "INVALID_ROLE", handCfg.JwtSecret)
		usrSvc.EXPECT().GetUserById("INVALID_ID").Return(response.User{}, config.SvcInvalidToken, errors.New("Invalid token"))

		req := meRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		mockToken := ""
		req := meRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestGetAllUsers(t *testing.T) {
	getAllUsersRequest := func(t *testing.T, page int, search, token string) *http.Request {
		url := fmt.Sprintf("/api/users?page=%v&search=%v", page, search)
		req, err := http.NewRequest("GET", url, nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on success", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockUser := response.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role, handCfg.JwtSecret)
		page, search := 1, ""
		usrSvc.EXPECT().GetAllUsers(page, search).Return([]response.User{mockUser}, config.SvcEmptyMsg, nil)

		req := getAllUsersRequest(t, page, search, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		mockToken := ""
		req := getAllUsersRequest(t, 1, "", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestGetUsersById(t *testing.T) {
	getUsersByIdRequest := func(t *testing.T, id, token string) *http.Request {
		req, err := http.NewRequest("GET", "/api/users/"+id, nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on success", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockUser := response.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role, handCfg.JwtSecret)
		usrSvc.EXPECT().GetUserById(mockUser.ID).Return(mockUser, config.SvcEmptyMsg, nil)

		req := getUsersByIdRequest(t, mockUser.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockUser := response.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role, handCfg.JwtSecret)
		usrSvc.EXPECT().GetUserById("INVALID_ID").Return(response.User{}, config.SvcInvalidToken, errors.New("Invalid id"))

		req := getUsersByIdRequest(t, "INVALID_ID", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		mockToken := ""

		req := getUsersByIdRequest(t, "some_id", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestCreateUser(t *testing.T) {
	createUserRequest := func(t *testing.T, body, token string) *http.Request {
		req, err := http.NewRequest("POST", "/api/users", bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 201 on success", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockAdmin := entity.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role, handCfg.JwtSecret)

		mockUser := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		usrSvc.EXPECT().CreateUser(mockUser).Return(mockAdmin.ToResponse(), config.SvcUserCreated, nil)

		body, _ := json.Marshal(mockUser)
		req := createUserRequest(t, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("should return 403 on non-admin", func(t *testing.T) {
		r, _, handCfg := makeUserService(t)

		mockNonAdmin := entity.User{ID: "non-admin", Email: "non-admin", Password: "non-admin", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockNonAdmin.ID, mockNonAdmin.Role, handCfg.JwtSecret)

		req := createUserRequest(t, "", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on user already exists", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockAdmin := entity.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role, handCfg.JwtSecret)

		mockUser := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		usrSvc.EXPECT().CreateUser(mockUser).Return(response.User{}, config.SvcUserExists, errors.New("User already exists"))

		body, _ := json.Marshal(mockUser)
		req := createUserRequest(t, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		mockToken := ""

		req := createUserRequest(t, "", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	updateUserRequest := func(t *testing.T, id, body, token string) *http.Request {
		req, err := http.NewRequest("PATCH", "/api/users/"+id, bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on update by owner", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockOwner := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockOwner.ID, mockOwner.Role, handCfg.JwtSecret)

		updateBody := request.User{Email: "NEW"}
		updatedUser := response.User{ID: "user", Email: "NEW", Password: "user", Role: "user"}
		usrSvc.EXPECT().UpdateUser(mockOwner.ID, updateBody).Return(updatedUser, config.SvcUserUpdated, nil)

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, mockOwner.ID, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 200 on update by admin", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockAdmin := request.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role, handCfg.JwtSecret)

		mockUser := entity.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		updateBody := request.User{Email: "NEW"}
		updatedUser := response.User{ID: "user", Email: "NEW", Password: "user", Role: "user"}
		usrSvc.EXPECT().UpdateUser(mockUser.ID, updateBody).Return(updatedUser, config.SvcUserUpdated, nil)

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, mockUser.ID, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 403 on update by non-owner and non-admin", func(t *testing.T) {
		r, _, handCfg := makeUserService(t)

		mockUser1 := request.User{ID: "user1", Email: "user1", Password: "user1", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser1.ID, mockUser1.Role, handCfg.JwtSecret)

		mockUser2 := entity.User{ID: "user2", Email: "user2", Password: "user2", Role: "user"}
		updateBody := request.User{Email: "NEW"}

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, mockUser2.ID, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockAdmin := entity.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role, handCfg.JwtSecret)

		updateBody := request.User{Email: "NEW"}
		usrSvc.EXPECT().UpdateUser("INVALID_ID", updateBody).Return(response.User{}, config.SvcFailedUpdateUser, errors.New("No such user"))

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, "INVALID_ID", string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		mockToken := ""

		req := updateUserRequest(t, "ID", "BODY", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestDeleteUser(t *testing.T) {
	deleteUserRequest := func(t *testing.T, id, token string) *http.Request {
		req, err := http.NewRequest("DELETE", "/api/users/"+id, nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on delete by owner", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockOwner := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockOwner.ID, mockOwner.Role, handCfg.JwtSecret)

		usrSvc.EXPECT().DeleteUser(mockOwner.ID).Return(response.User(mockOwner.ToEntity()), config.SvcUserDeleted, nil)

		req := deleteUserRequest(t, mockOwner.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 200 on delete by admin", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockUser := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		mockAdmin := request.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role, handCfg.JwtSecret)

		usrSvc.EXPECT().DeleteUser(mockUser.ID).Return(response.User(mockUser.ToEntity()), config.SvcUserDeleted, nil)

		req := deleteUserRequest(t, mockUser.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 403 on delete by non-owner and non-admin", func(t *testing.T) {
		r, _, handCfg := makeUserService(t)

		mockUser1 := request.User{ID: "user1", Email: "user1", Password: "user1", Role: "user"}
		mockUser2 := request.User{ID: "user2", Email: "user2", Password: "user2", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser1.ID, mockUser1.Role, handCfg.JwtSecret)

		req := deleteUserRequest(t, mockUser2.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r, usrSvc, handCfg := makeUserService(t)

		mockAdmin := request.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role, handCfg.JwtSecret)

		usrSvc.EXPECT().DeleteUser("INVALID_ID").Return(response.User{}, config.SvcFailedDeleteUser, errors.New("No such user"))

		req := deleteUserRequest(t, "INVALID_ID", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _, _ := makeUserService(t)

		mockToken := ""

		req := deleteUserRequest(t, "ID", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}
