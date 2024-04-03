package service

import (
	"sigma-test/config"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
	"sigma-test/internal/util"
)

type UserConfig struct {
	JwtSecret string
}

type UserService struct {
	repo UserRepo
	cfg  UserConfig
}

func NewUserService(r UserRepo, c UserConfig) UserService {
	return UserService{repo: r, cfg: c}
}

func (s UserService) SignUp(body request.User) (response.User, error) {
	_, err := s.GetUserByEmail(body.Email)
	if err == nil {
		return response.User{}, config.ErrUserExists
	}

	return s.CreateUser(request.User{Email: body.Email, Password: body.Password, Role: body.Role})
}

func (s UserService) Login(body request.User) (string, error) {
	user, err := s.GetUserByEmail(body.Email)
	if err != nil {
		return "", config.ErrInvalidCredentials
	}

	if _, err := util.ComparePasswordAndHash(body.Password, user.Password); err != nil {
		return "", config.ErrInvalidCredentials
	}

	return util.GenerateJWTToken(user.ID, user.Role, s.cfg.JwtSecret)
}

func (s UserService) GetAllUsers(page int, search string) ([]response.User, error) {
	users, err := s.repo.GetUsers(page, search)
	if err != nil {
		return []response.User{}, config.ErrFailedGetUser
	}

	responses := make([]response.User, len(users))
	for i, v := range users {
		responses[i] = v.ToResponse()
	}

	return responses, nil
}

func (s UserService) GetUserById(id string) (response.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return response.User{}, config.ErrNoUser
	}

	return user.ToResponse(), nil
}

func (s UserService) GetUserByEmail(email string) (response.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return response.User{}, config.ErrNoUser
	}

	return user.ToResponse(), nil
}

func (s UserService) CreateUser(data request.User) (response.User, error) {
	hash, err := util.HashPassword(data.Password)
	if err != nil {
		return response.User{}, config.ErrFailedHashPassword
	}

	user := data.ToEntity()
	user.Password = hash

	new_user, err := s.repo.CreateUser(user)
	if err != nil {
		return response.User{}, config.ErrFailedCreateUser
	}

	return new_user.ToResponse(), nil
}

func (s UserService) UpdateUser(id string, data request.User) (response.User, error) {
	hash, err := util.HashPassword(data.Password)
	if err != nil {
		return response.User{}, config.ErrFailedHashPassword
	}

	updateBody := data.ToEntity()
	if data.Password != "" {
		updateBody.Password = hash
	}

	updatedUser, err := s.repo.UpdateUser(id, updateBody)
	if err != nil {
		return response.User{}, config.ErrFailedUpdateUser
	}

	return updatedUser.ToResponse(), nil
}

func (s UserService) DeleteUser(id string) (response.User, error) {
	user, err := s.repo.DeleteUser(id)
	if err != nil {
		return response.User{}, config.ErrFailedDeleteUser
	}

	return user.ToResponse(), nil
}
