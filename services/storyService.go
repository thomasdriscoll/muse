package services

import (
	"errors"

	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
	"github.com/thomasdriscoll/muse/repositories"
)

type StoryService interface {
	GetStoryById(ID uint64) (*models.Story, error)
	CreateStory(storyRequest *models.StoryFromURLRequest) (*models.Story, error)
	DeleteById(ID uint64) error
	GetStoryByRandom() (*models.Story, error)
	GetStoriesByAuthorId(authorId uint64) (*[]models.Story, error)
	GetStoriesByTag(tag string) (*[]models.Story, error)
}

type StoryServiceImpl struct {
	StoryRepo repositories.StoryRepository
}

// func scrape(url string, urlType string) (*models.Story, error) {
// 	return nil, nil
// }

func (s *StoryServiceImpl) GetStoryById(ID uint64) (*models.Story, error) {
	var storyFromID *models.Story
	storyFromID, err := s.StoryRepo.GetStoryById(ID)
	if err != nil {
		// Add logger statement here
		return nil, errors.New(enums.ErrorStoryNotFound)
	}
	return storyFromID, nil
}

func (s *StoryServiceImpl) CreateStory(storyRequest *models.StoryFromURLRequest) (*models.Story, error) {
	return nil, nil
}

func (s *StoryServiceImpl) DeleteById(ID uint64) error {
	return nil
}
func (s *StoryServiceImpl) GetStoryByRandom() (*models.Story, error) {
	return nil, nil
}
func (s *StoryServiceImpl) GetStoriesByAuthorId(authorId uint64) (*[]models.Story, error) {
	return nil, nil
}

func (s *StoryServiceImpl) GetStoriesByTag(tag string) (*[]models.Story, error) {
	return nil, nil
}
