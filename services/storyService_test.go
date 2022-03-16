package services

import (
	"testing"

	"github.com/thomasdriscoll/muse/models"
)

// Global variables
var storyId string = "goods3bucketid"
var notFoundId string = "notfounds3bucketId"
var badId string = "400id"

func TestGetStoryById(t *testing.T) {

	testCases := []struct {
		testMessage   string
		id            string
		expectedStory models.Story
	}{
		{
			testMessage:   "Happy path for StoryService.GetStoryById",
			id:            storyId,
			expectedStory: models.Story{},
		},
		{
			testMessage:   "Not found path for StoryService.GetStoryById",
			id:            notFoundId,
			expectedStory: models.Story{},
		},
		{
			testMessage:   "Bad request path for StoryService.GetStoryById",
			id:            badId,
			expectedStory: models.Story{},
		},
	}

	for _, testCase := range testCases {
		t.Run(test)
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
