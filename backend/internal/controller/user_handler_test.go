package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigma-test/internal/controller/mocks"
	"sigma-test/internal/entity"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
	"sigma-test/internal/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeUserService(t *testing.T) (*gin.Engine, *mocks.UserService) {
	r := gin.New()
	userService := mocks.NewUserService(t)
	InitUserHandler(r, userService)

	return r, userService
}

func TestSignup(t *testing.T) {
	signUpRequest := func(t *testing.T, body string) *http.Request {
		req, err := http.NewRequest("POST", "/api/user/signup", bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		return req
	}
	t.Run("should return 200 on successful signup", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test123", Role: "user"}
		usrSvc.EXPECT().SignUp(mockUser).Return(response.User(mockUser), nil)

		body, _ := json.Marshal(mockUser)
		req := signUpRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r, _ := makeUserService(t)

		req := signUpRequest(t, "")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 if user exists", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test123", Role: "user"}
		usrSvc.EXPECT().SignUp(mockUser).Return(response.User{}, errors.New("User already exists"))

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
		r, usrSvc := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		usrSvc.EXPECT().Login(mockUser).Return("token", nil)

		body, _ := json.Marshal(mockUser)
		req := loginRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 400 on invalid body", func(t *testing.T) {
		r, _ := makeUserService(t)

		req := loginRequest(t, "")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return 422 on invalid credantials", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := request.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		usrSvc.EXPECT().Login(mockUser).Return("", errors.New("Invalid Credantials"))

		body, _ := json.Marshal(mockUser)
		req := loginRequest(t, string(body))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
}

func TestMe(t *testing.T) {
	meRequest := func(t *testing.T, token string) *http.Request {
		req, err := http.NewRequest("GET", "/api/user/me", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on success", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := entity.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)
		usrSvc.EXPECT().GetUserById(mockUser.ID).Return(mockUser.ToResponse(), nil)

		req := meRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid token", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockToken, _ := util.GenerateJWTToken("INVALID_ID", "INVALID_ROLE")
		usrSvc.EXPECT().GetUserById("INVALID_ID").Return(response.User{}, errors.New("Invalid token"))

		req := meRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockToken := ""
		req := meRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestGetAllUsers(t *testing.T) {
	getAllUsersRequest := func(t *testing.T, token string) *http.Request {
		req, err := http.NewRequest("GET", "/api/users", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on success", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := response.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)
		usrSvc.EXPECT().GetAllUsers().Return([]response.User{mockUser}, nil)

		req := getAllUsersRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockToken := ""
		req := getAllUsersRequest(t, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestGetUsersById(t *testing.T) {
	getUsersByIdRequest := func(t *testing.T, id, token string) *http.Request {
		req, err := http.NewRequest("GET", "/api/user?id="+id, nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on success", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := response.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)
		usrSvc.EXPECT().GetUserById(mockUser.ID).Return(mockUser, nil)

		req := getUsersByIdRequest(t, mockUser.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := response.User{ID: "test", Email: "test@test.com", Password: "test", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser.ID, mockUser.Role)
		usrSvc.EXPECT().GetUserById("INVALID_ID").Return(response.User{}, errors.New("Invalid id"))

		req := getUsersByIdRequest(t, "INVALID_ID", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockToken := ""

		req := getUsersByIdRequest(t, "", mockToken)
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
		r, usrSvc := makeUserService(t)

		mockAdmin := entity.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		mockUser := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		usrSvc.EXPECT().CreateUser(mockUser).Return(response.User(mockUser), nil)

		body, _ := json.Marshal(mockUser)
		req := createUserRequest(t, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("should return 403 on non-admin", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockNonAdmin := entity.User{ID: "non-admin", Email: "non-admin", Password: "non-admin", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockNonAdmin.ID, mockNonAdmin.Role)

		req := createUserRequest(t, "", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on user already exists", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockAdmin := entity.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		mockUser := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		usrSvc.EXPECT().CreateUser(mockUser).Return(response.User{}, errors.New("User already exists"))

		body, _ := json.Marshal(mockUser)
		req := createUserRequest(t, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockToken := ""

		req := createUserRequest(t, "", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	updateUserRequest := func(t *testing.T, id, body, token string) *http.Request {
		req, err := http.NewRequest("PATCH", "/api/users?id="+id, bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on update by owner", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockOwner := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockOwner.ID, mockOwner.Role)

		updateBody := request.User{Email: "NEW"}
		updatedUser := response.User{ID: "user", Email: "NEW", Password: "user", Role: "user"}
		usrSvc.EXPECT().UpdateUser(mockOwner.ID, updateBody).Return(updatedUser, nil)

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, mockOwner.ID, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 200 on update by admin", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockAdmin := request.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		mockUser := entity.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		updateBody := request.User{Email: "NEW"}
		updatedUser := response.User{ID: "user", Email: "NEW", Password: "user", Role: "user"}
		usrSvc.EXPECT().UpdateUser(mockUser.ID, updateBody).Return(updatedUser, nil)

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, mockUser.ID, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 403 on update by non-owner and non-admin", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockUser1 := request.User{ID: "user1", Email: "user1", Password: "user1", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser1.ID, mockUser1.Role)

		mockUser2 := entity.User{ID: "user2", Email: "user2", Password: "user2", Role: "user"}
		updateBody := request.User{Email: "NEW"}

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, mockUser2.ID, string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockAdmin := entity.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		updateBody := request.User{Email: "NEW"}
		usrSvc.EXPECT().UpdateUser("INVALID_ID", updateBody).Return(response.User{}, errors.New("No such user"))

		body, _ := json.Marshal(updateBody)
		req := updateUserRequest(t, "INVALID_ID", string(body), mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockToken := ""

		req := updateUserRequest(t, "ID", "BODY", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func TestDeleteUser(t *testing.T) {
	deleteUserRequest := func(t *testing.T, id, token string) *http.Request {
		req, err := http.NewRequest("DELETE", "/api/users?id="+id, nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		return req
	}
	t.Run("should return 200 on delete by owner", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockOwner := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockOwner.ID, mockOwner.Role)

		usrSvc.EXPECT().DeleteUser(mockOwner.ID).Return(nil)

		req := deleteUserRequest(t, mockOwner.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 200 on delete by admin", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockUser := request.User{ID: "user", Email: "user", Password: "user", Role: "user"}
		mockAdmin := request.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		usrSvc.EXPECT().DeleteUser(mockUser.ID).Return(nil)

		req := deleteUserRequest(t, mockUser.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("should return 403 on delete by non-owner and non-admin", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockUser1 := request.User{ID: "user1", Email: "user1", Password: "user1", Role: "user"}
		mockUser2 := request.User{ID: "user2", Email: "user2", Password: "user2", Role: "user"}
		mockToken, _ := util.GenerateJWTToken(mockUser1.ID, mockUser1.Role)

		req := deleteUserRequest(t, mockUser2.ID, mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusForbidden, res.Code)
	})

	t.Run("should return 422 on invalid id", func(t *testing.T) {
		r, usrSvc := makeUserService(t)

		mockAdmin := request.User{ID: "admin", Email: "admin", Password: "admin", Role: "admin"}
		mockToken, _ := util.GenerateJWTToken(mockAdmin.ID, mockAdmin.Role)

		usrSvc.EXPECT().DeleteUser("INVALID_ID").Return(errors.New("No such user"))

		req := deleteUserRequest(t, "INVALID_ID", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("should return 401 on missing token", func(t *testing.T) {
		r, _ := makeUserService(t)

		mockToken := ""

		req := deleteUserRequest(t, "ID", mockToken)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}
