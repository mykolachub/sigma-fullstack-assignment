package service

import (
	"sigma-user/config"
	"sigma-user/internal/request"
	"sigma-user/internal/response"
	"sigma-user/internal/util"
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

func (s UserService) SignUp(body request.User) (response.User, config.ServiceCode, error) {
	_, _, err := s.GetUserByEmail(body.Email)
	if err == nil {
		return response.User{}, config.SvcUserExists, config.SvcUserExists.ToError()

	}

	return s.CreateUser(request.User{Email: body.Email, Password: body.Password, Role: body.Role})
}

func (s UserService) Login(body request.User) (string, config.ServiceCode, error) {
	user, _, err := s.GetUserByEmail(body.Email)
	if err != nil {
		return "", config.SvcInvalidCredentials, config.SvcInvalidCredentials.ToError()
	}

	if _, err := util.ComparePasswordAndHash(body.Password, user.Password); err != nil {
		return "", config.SvcInvalidCredentials, config.SvcInvalidCredentials.ToError()
	}

	token, err := util.GenerateJWTToken(user.ID, user.Role, s.cfg.JwtSecret)
	if err != nil {
		return "", config.SvcFailedCreateToken, config.SvcFailedCreateToken.ToError()
	}

	return token, config.SvcEmptyMsg, nil
}

func (s UserService) GetAllUsers(page int, search string) ([]response.User, config.ServiceCode, error) {
	users, err := s.repo.GetUsers(page, search)
	if err != nil {
		return []response.User{}, config.SvcFailedGetUser, config.SvcFailedGetUser.ToError()
	}

	responses := make([]response.User, len(users))
	for i, v := range users {
		responses[i] = v.ToResponse()
	}

	return responses, config.SvcEmptyMsg, nil
}

func (s UserService) GetUserById(id string) (response.User, config.ServiceCode, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return response.User{}, config.SvcNoUser, config.SvcNoUser.ToError()
	}

	return user.ToResponse(), config.SvcEmptyMsg, nil
}

func (s UserService) GetUserByEmail(email string) (response.User, config.ServiceCode, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return response.User{}, config.SvcNoUser, config.SvcNoUser.ToError()
	}

	return user.ToResponse(), config.SvcEmptyMsg, nil
}

func (s UserService) CreateUser(data request.User) (response.User, config.ServiceCode, error) {
	hash, err := util.HashPassword(data.Password)
	if err != nil {
		return response.User{}, config.SvcFailedHashPassword, config.SvcFailedHashPassword.ToError()
	}

	user := data.ToEntity()
	user.Password = hash

	new_user, err := s.repo.CreateUser(user)
	if err != nil {
		return response.User{}, config.SvcFailedCreateUser, config.SvcFailedCreateUser.ToError()
	}

	return new_user.ToResponse(), config.SvcUserCreated, nil
}

func (s UserService) UpdateUser(id string, data request.User) (response.User, config.ServiceCode, error) {
	hash, err := util.HashPassword(data.Password)
	if err != nil {
		return response.User{}, config.SvcFailedHashPassword, config.SvcFailedHashPassword.ToError()
	}

	updateBody := data.ToEntity()
	if data.Password != "" {
		updateBody.Password = hash
	}

	updatedUser, err := s.repo.UpdateUser(id, updateBody)
	if err != nil {
		return response.User{}, config.SvcFailedUpdateUser, config.SvcFailedUpdateUser.ToError()
	}

	return updatedUser.ToResponse(), config.SvcUserUpdated, nil
}

func (s UserService) DeleteUser(id string) (response.User, config.ServiceCode, error) {
	user, err := s.repo.DeleteUser(id)
	if err != nil {
		return response.User{}, config.SvcFailedDeleteUser, config.SvcFailedDeleteUser.ToError()
	}

	return user.ToResponse(), config.SvcUserDeleted, nil
}
