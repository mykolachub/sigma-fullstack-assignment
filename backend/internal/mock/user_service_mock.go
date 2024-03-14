package mock

import (
	"errors"
	"log"
	"sigma-test/internal/entity"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
)

type MockUserService struct {
	MockDB []entity.User
}

func (m *MockUserService) DeleteUser(email string) error {
	return nil
}

func (m *MockUserService) GetAllUsers() ([]response.User, error) {
	var users []response.User
	for _, v := range m.MockDB {
		users = append(users, v.ToResponse())
	}
	return users, nil
}

func (m *MockUserService) GetUserById(id string) (response.User, error) {
	for _, v := range m.MockDB {
		if v.ID == id {
			return v.ToResponse(), nil
		}
	}

	return response.User{}, errors.New("no such user")
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

func (m *MockUserService) Login(body request.User) (string, error) {
	// If user exists in mockDB, return test access token
	for _, v := range m.MockDB {
		if v.Email == body.Email {
			return "testTokenStrign", nil
		}
	}

	return "", errors.New("test login error")
}
