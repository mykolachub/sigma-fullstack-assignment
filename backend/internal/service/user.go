package service

import (
	"errors"
	"fmt"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
	"sigma-test/internal/util"
	"sigma-test/pkg/helpers"

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
	users, err := s.repo.GetAllUsers()

	responses := make([]response.User, len(users))
	for i, v := range users {
		responses[i] = v.ToResponse()
	}

	return responses, err
}

func (s UserService) GetUserById(id string) (response.User, error) {
	idx, exists := s.repo.GetUserIdxById(id)
	if !exists {
		return response.User{}, errors.New("No such user")
	}

	user, err := s.repo.GetUserByIdx(idx)
	return user.ToResponse(), err
}

func (s UserService) GetUserByEmail(email string) (response.User, error) {
	idx, exists := s.repo.GetUserIdxByEmail(email)
	if !exists {
		return response.User{}, errors.New("No such user")
	}

	user, err := s.repo.GetUserByIdx(idx)
	return user.ToResponse(), err
}

func (s UserService) CreateUser(user request.User) (response.User, error) {
	_, exists := s.repo.GetUserIdxByEmail(user.Email)
	if exists {
		return response.User{}, errors.New("User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.User{}, errors.New("failed to hash body")
	}

	userEntity := user.ToEntity()
	userEntity.ID = helpers.GetKsuid()
	userEntity.Password = string(hash)

	new_user, err := s.repo.AddUser(userEntity)
	return new_user.ToResponse(), err
}

func (s UserService) UpdateUser(id string, body request.User) (response.User, error) {
	idx, exists := s.repo.GetUserIdxById(id)
	if !exists {
		return response.User{}, errors.New("No such user")
	}

	fmt.Println(body)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return response.User{}, errors.New("failed to hash body")
	}

	bodyEntity := body.ToEntity()
	bodyEntity.ID = id
	bodyEntity.Password = string(hash)

	updatedUser, err := s.repo.UpdateUser(bodyEntity, idx)
	return updatedUser.ToResponse(), err
}

func (s UserService) DeleteUser(id string) error {
	idx, exists := s.repo.GetUserIdxById(id)
	if !exists {
		return errors.New("No such user")
	}

	err := s.repo.DeleteUser(idx)
	return err
}
