package service

import (
	"errors"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
	"sigma-test/pkg/helpers"
)

type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) UserService {
	return UserService{repo: r}
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

func (s UserService) CreateUser(user request.User) (response.User, error) {
	_, exists := s.repo.GetUserIdxByEmail(user.Email)
	if exists {
		return response.User{}, errors.New("User already exists")
	}

	userEntity := user.ToEntity()
	userEntity.ID = helpers.GetKsuid()

	new_user, err := s.repo.AddUser(userEntity)
	return new_user.ToResponse(), err
}

func (s UserService) UpdateUser(id string, body request.User) (response.User, error) {
	idx, exists := s.repo.GetUserIdxById(id)
	if !exists {
		return response.User{}, errors.New("No such user")
	}

	bodyEntity := body.ToEntity()
	bodyEntity.ID = id

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
