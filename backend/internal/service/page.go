package service

import "sigma-test/internal/response"

type PageService struct {
	repo PageRepo
}

func NewPageService(r PageRepo) PageService {
	return PageService{repo: r}
}

func (s PageService) TrackPage(name string) error {
	return s.repo.TrackPage(name)
}

func (s PageService) GetPageCount(name string) (response.Page, error) {
	page, err := s.repo.GetPage(name)
	return response.Page(page), err
}
