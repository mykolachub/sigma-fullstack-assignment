package service

import (
	"errors"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
	"sigma-test/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) UserService {
	return UserService{repo: r}
}

func (s UserService) SignUp(body request.User) (response.User, error) {
	_, err := s.GetUserByEmail(body.Email)
	if err == nil {
		return response.User{}, errors.New("user already exists")
	}

	return s.CreateUser(request.User{Email: body.Email, Password: body.Password, Role: body.Role})
}

func (s UserService) Login(body request.User) (string, error) {
	user, err := s.GetUserByEmail(body.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return "", err
	}

	return util.GenerateJWTToken(user.ID, user.Role)
}

func (s UserService) GetAllUsers() ([]response.User, error) {
	users, err := s.repo.GetUsers()

	responses := make([]response.User, len(users))
	for i, v := range users {
		responses[i] = v.ToResponse()
	}

	return responses, err
}

func (s UserService) GetUserById(id string) (response.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return response.User{}, errors.New("No such user")
	}

	return user.ToResponse(), nil
}

func (s UserService) GetUserByEmail(email string) (response.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return response.User{}, errors.New("No such user")
	}

	return user.ToResponse(), nil
}

func (s UserService) CreateUser(data request.User) (response.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return response.User{}, errors.New("failed to hash body")
	}

	user := data.ToEntity()
	user.Password = string(hash)

	new_user, err := s.repo.CreateUser(user)
	if err != nil {
		return response.User{}, errors.New("failed to create user")

	}
	return new_user.ToResponse(), nil
}

func (s UserService) UpdateUser(id string, data request.User) (response.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return response.User{}, errors.New("failed to hash body")
	}

	updateBody := data.ToEntity()
	if data.Password != "" {
		updateBody.Password = string(hash)
	}

	updatedUser, err := s.repo.UpdateUser(id, updateBody)
	return updatedUser.ToResponse(), err
}

func (s UserService) DeleteUser(id string) (response.User, error) {
	user, err := s.repo.DeleteUser(id)
	return user.ToResponse(), err
}
