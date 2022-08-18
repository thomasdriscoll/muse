package services

import (
	"github.com/thomasdriscoll/muse/models"
	"github.com/thomasdriscoll/muse/repositories"
)

type StoryService interface {
	GetStoryById(ID string) (*models.Story, error)
	CreateStory(storyRequest *models.StoryFromURLRequest) (*models.Story, error)
	DeleteById(ID string) error
	GetStoryByRandom() (*models.Story, error)
	GetStoriesByAuthorId(authorId string) (*[]models.Story, error)
	GetStoriesByTag(tag string) (*[]models.Story, error)
}

type StoryServiceImpl struct {
	Scrapper          Scrapper
	StoryMetadataRepo repositories.StoryMetadataRepository
}

func (s *StoryServiceImpl) GetStoryById(ID string) (*models.Story, error) {
	storyMetadataFromID, err := s.StoryMetadataRepo.GetStoryById(ID)
	if err != nil {
		// Add logger statement here
		return nil, err
	}
	storyContent, err := s.Scrapper.Scrape(storyMetadataFromID.Source, storyMetadataFromID.SourceType)
	if err != nil {
		return nil, err
	}
	storyFromID := models.Story{
		StoryMetadata: *storyMetadataFromID,
		Content:       storyContent,
	}
	return &storyFromID, nil
}

func (s *StoryServiceImpl) CreateStory(storyRequest *models.StoryFromURLRequest) (*models.Story, error) {
	return nil, nil
}

func (s *StoryServiceImpl) DeleteById(ID string) error {
	return nil
}
func (s *StoryServiceImpl) GetStoryByRandom() (*models.Story, error) {
	return nil, nil
}
func (s *StoryServiceImpl) GetStoriesByAuthorId(authorId string) (*[]models.Story, error) {
	return nil, nil
}

func (s *StoryServiceImpl) GetStoriesByTag(tag string) (*[]models.Story, error) {
	return nil, nil
}
