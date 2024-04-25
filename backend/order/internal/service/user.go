package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sigma-order/internal/entity"
)

type UserServiceConfig struct {
	Url string
}

type UserService struct {
	cfg UserServiceConfig
}

func NewUserService(config UserServiceConfig) UserService {
	return UserService{cfg: config}
}

func (s *UserService) GetUser(ctx context.Context, userId string) (*entity.User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", s.cfg.Url+userId, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Code int `json:"code"`
		Data struct {
			User entity.User `json:"user"`
		} `json:"data"`
		Status string `json:"status"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data.User, nil
}

func (s *UserService) UserExists(ctx context.Context, userId string) (bool, error) {
	user, err := s.GetUser(ctx, userId)
	if err != nil {
		return false, err
	}

	if user.ID == "" {
		return false, nil
	}

	return true, nil
}
