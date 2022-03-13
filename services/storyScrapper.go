package services

import "github.com/thomasdriscoll/muse/models"

type StoryScrapper interface {
	Scrape(url string, urlType string) (*models.Story, error)
}

type StoryScrapperImpl struct {
}

func (s *StoryScrapperImpl) Scrape(url string, urlType string) (*models.Story, error) {
	return nil, nil
}
