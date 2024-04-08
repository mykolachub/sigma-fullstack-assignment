package service

import (
	"sigma-test/config"
	"sigma-test/internal/response"
)

type PageService struct {
	repo PageRepo
}

func NewPageService(r PageRepo) PageService {
	return PageService{repo: r}
}

func (s PageService) TrackPage(name string) (config.ServiceCode, error) {
	err := s.repo.TrackPage(name)
	if err != nil {
		return config.SvcFailedTrackPage, config.SvcFailedTrackPage.ToError()
	}

	return config.SvcPageTracked, nil
}

func (s PageService) GetPageCount(name string) (response.Page, config.ServiceCode, error) {
	page, err := s.repo.GetPage(name)
	if err != nil {
		return response.Page{}, config.SvcFailedGetPage, config.SvcFailedGetPage.ToError()
	}
	return response.Page(page), config.SvcEmptyMsg, nil
}
