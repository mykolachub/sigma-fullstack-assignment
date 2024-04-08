package service

import (
	"sigma-test/config"
	"sigma-test/internal/response"

	"github.com/adrianbrad/queue"
)

type PageService struct {
	repo PageRepo
}

func NewPageService(r PageRepo) PageService {
	return PageService{repo: r}
}

func (s PageService) TrackPage(q *queue.Linked[string], name string) (config.ServiceCode, error) {
	if q.Size() > config.PageMaxQueueSize {
		err := s.repo.BatchTrackPages(q)
		if err != nil {
			return config.SvcFailedTrackPage, config.SvcFailedTrackPage.ToError()
		}
		q.Clear()
	}
	q.Offer(name)

	return config.SvcPageTracked, nil
}

func (s PageService) GetPageCount(name string) (response.Page, config.ServiceCode, error) {
	page, err := s.repo.GetPage(name)
	if err != nil {
		return response.Page{}, config.SvcFailedGetPage, config.SvcFailedGetPage.ToError()
	}
	return response.Page(page), config.SvcEmptyMsg, nil
}
