package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
	"github.com/thomasdriscoll/muse/testhelper"
)

// Global variables
var storyId string = "goods3bucketid"
var storyNotFoundId string = "notfounds3bucketId"
var badId string = "dangerWillRobinson!"

var globalAuthorId string = "authorId"
var authorNotFoundId string = "authorNotFoundId"

var globalTag string = "superhero"
var tagNotFound string = "super-niche-content"

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
			id:            storyNotFoundId,
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
			if reflect.DeepEqual(*testCase.expectedStory, *story) {
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

	testCases := []struct {
		testMessage     string
		authorId        string
		expectedStories *[]models.Story
		err             error
	}{
		{
			testMessage: "Happy path for StoryService.GetStoryByAuthorId",
			authorId:    globalAuthorId,
			expectedStories: &[]models.Story{
				{
					StoryMetadata: testhelper.GetStoryMetadata(),
					Content:       testhelper.GetStoryContent(),
				},
			},
			err: nil,
		},
		{
			testMessage:     "Not found path for StoryService.GetStoryByAuthorId",
			authorId:        authorNotFoundId,
			expectedStories: nil,
			err:             errors.New(enums.ErrorAuthorNotFound),
		},
		{
			testMessage:     "Bad request path for StoryService.GetStoryByAuthorId",
			authorId:        badId,
			expectedStories: nil,
			err:             errors.New(enums.ErrorInvalidStoryRequest),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			story, err := storyService.GetStoriesByAuthorId(testCase.authorId)
			if reflect.DeepEqual((*testCase.expectedStories)[0], *story) {
				t.Errorf("Expected story does not match actual story in GetStoryById")
			}
			if testCase.err != err {
				t.Errorf("Expected error does not match actual error in GetStoryById")
			}
		})
	}

}

func TestGetStoriesByTag(t *testing.T) {
	testCases := []struct {
		testMessage     string
		tag             string
		expectedStories *[]models.Story
		err             error
	}{
		{
			testMessage: "Happy path for StoryService.GetStoryByAuthorId",
			tag:         globalTag,
			expectedStories: &[]models.Story{
				{
					StoryMetadata: testhelper.GetStoryMetadata(),
					Content:       testhelper.GetStoryContent(),
				},
			},
			err: nil,
		},
		{
			testMessage:     "Not found path for StoryService.GetStoryByAuthorId",
			tag:             tagNotFound,
			expectedStories: nil,
			err:             errors.New(enums.ErrorAuthorNotFound),
		},
		{
			testMessage:     "Bad request path for StoryService.GetStoryByAuthorId",
			tag:             badId,
			expectedStories: nil,
			err:             errors.New(enums.ErrorInvalidStoryRequest),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			story, err := storyService.GetStoriesByTag(testCase.tag)
			if reflect.DeepEqual((*testCase.expectedStories)[0], *story) {
				t.Errorf("Expected story does not match actual story in GetStoryById")
			}
			if testCase.err != err {
				t.Errorf("Expected error does not match actual error in GetStoryById")
			}
		})
	}
}

// *************************************************************************************************************
// 				MOCKS
// *************************************************************************************************************

type MockScrapper struct{}

type MockStoryMetadataRepo struct{}

func (mr *MockStoryMetadataRepo) GetStoryById(ID string) (*models.StoryMetadata, error) {
	if ID == storyId {
		metadata := testhelper.GetStoryMetadata()
		return &metadata, nil
	}
	return nil, errors.New(enums.ErrorStoryNotFound)
}
func (mr *MockStoryMetadataRepo) Save(story *models.StoryMetadata) error {
	if story.StoryId == storyId {
		return nil
	} else if story.StoryId == badId {
		return errors.New(enums.ErrorStoryAlreadyExists)
	}
	return nil
}
func (mr *MockStoryMetadataRepo) DeleteById(ID string) error {
	if ID == storyId {
		return nil
	}
	return errors.New(enums.ErrorStoryNotFound)

}
func (mr *MockStoryMetadataRepo) GetStoryByRandom() (*models.StoryMetadata, error) {
	metadata := testhelper.GetStoryMetadata()
	return &metadata, nil
}

func (mr *MockStoryMetadataRepo) GetStoriesByAuthorId(authorId string) (*[]models.StoryMetadata, error) {
	if authorId == globalAuthorId {
		metadata := testhelper.GetStoryMetadata()
		listOfMetadata := []models.StoryMetadata{metadata}
		return &listOfMetadata, nil
	}
	return nil, errors.New(enums.ErrorAuthorNotFound)
}
func (mr *MockStoryMetadataRepo) GetStoriesByTag(tag string) (*[]models.StoryMetadata, error) {
	if tag == globalTag {
		metadata := testhelper.GetStoryMetadata()
		listOfMetadata := []models.StoryMetadata{metadata}
		return &listOfMetadata, nil
	}
	return nil, errors.New(enums.ErrorTagNotFound)

}
