package services

import (
	"errors"
	"testing"

	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
	"github.com/thomasdriscoll/muse/testhelper"
)

// Global variables
var storyId string = "goods3bucketid"
var notFoundId string = "notfounds3bucketId"
var badId string = "400id"
var storyService StoryServiceImpl = StoryServiceImpl{
	Scrapper:          &MockScrapper{},
	StoryMetadataRepo: &MockStoryMetadataRepo{},
}

// *************************************************************************************************************
// 				TESTS
// *************************************************************************************************************

func TestGetStoryById(t *testing.T) {

	testCases := []struct {
		testMessage   string
		id            string
		expectedStory *models.Story
		err           error
	}{
		{
			testMessage: "Happy path for StoryService.GetStoryById",
			id:          storyId,
			expectedStory: &models.Story{
				StoryMetadata: testhelper.GetStoryMetadata(),
				Content:       testhelper.GetStoryContent(),
			},
			err: nil,
		},
		{
			testMessage:   "Not found path for StoryService.GetStoryById",
			id:            notFoundId,
			expectedStory: nil,
			err:           errors.New(enums.ErrorStoryNotFound),
		},
		{
			testMessage:   "Bad request path for StoryService.GetStoryById",
			id:            badId,
			expectedStory: nil,
			err:           errors.New(enums.ErrorInvalidStoryId),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			story, err := storyService.GetStoryById(testCase.id)
			if &testCase.expectedStory != &story {
				t.Errorf("Expected story does not match actual story in GetStoryById")
			}
			if testCase.err != err {
				t.Errorf("Expected error does not match actual error in GetStoryById")
			}
		})
	}
}

func TestCreateStory(t *testing.T) {

}

func TestDeleteById(t *testing.T) {

}

func TestGetStoryByRandom(t *testing.T) {

}

func TestGetStoriesByAuthorId(t *testing.T) {

}

func TestGetStoriesByTag(t *testing.T) {

}

// *************************************************************************************************************
// 				MOCKS
// *************************************************************************************************************

type MockScrapper struct{}

type MockStoryMetadataRepo struct{}

func (mr *MockStoryMetadataRepo) GetStoryById(ID string) (*models.StoryMetadata, error) {
	return nil, nil
}
func (mr *MockStoryMetadataRepo) Save(story *models.StoryMetadata) error {
	return nil
}
func (mr *MockStoryMetadataRepo) DeleteById(ID string) error {
	return nil
}
func (mr *MockStoryMetadataRepo) GetStoryByRandom() (*models.StoryMetadata, error) {
	return nil, nil
}
func (mr *MockStoryMetadataRepo) GetStoriesByAuthorId(authorId string) (*[]models.StoryMetadata, error) {
	return nil, nil
}
func (mr *MockStoryMetadataRepo) GetStoriesByTag(tag string) (*[]models.StoryMetadata, error) {
	return nil, nil
}
