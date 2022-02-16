package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/thomasdriscoll/muse/models"
)

// Global variables
var storyId = 1 // 1 until I have something better
var storyController StoryController

type TestCase struct {
	writer               *httptest.ResponseRecorder
	request              *http.Request
	expectedResponseCode int
	expectedResponseBody []byte
	testMessage          string
}

func TestGetStory(t *testing.T) {
	// Constants
	route := "/story"
	story := models.Story{
		Metadata: models.StoryMetadata{},
		Content:  "content",
	}

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests & responses
	getRandomStoryRequest, _ := http.NewRequest(http.MethodGet, route, nil) // Add User Context here
	okResponse, _ := json.Marshal(story)
	notFoundResponse, _ := json.Marshal("Story not found")
	dbErrResponse, _ := json.Marshal("DB error")

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getRandomStoryRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(okResponse),
			testMessage:          "Happy path for StoryController.GetStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getRandomStoryRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Story not found for StoryController.GetStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getRandomStoryRequest,
			expectedResponseCode: http.StatusServiceUnavailable,
			expectedResponseBody: []byte(dbErrResponse),
			testMessage:          "Database error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.GetRandomStory(context)
			// assertions
			assert.Equal(t, testCase.expectedResponseBody, testCase.writer.Body.Bytes())
		})
	}
}

/*
func TestCreateStory(t *testing.T) {
	// Constants
	route := "/story"

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests
	createRequest, _ := http.NewRequest(http.MethodPost, route, nil) // Add User Context here

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              createRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(okResponse),
			testMessage:          "Happy path for StoryController.GetStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              okRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.CreateStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story request not valid for StoryController.CreateStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			context.Request = testCase.request
			storyController.CreateStory(context)
			// assertions

		})
	}

}

func TestGetStoryById(t *testing.T) {
	storyController := StoryControllerImpl{}
	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte("pong"),
			testMessage:          "Happy path for StoryController.GetStoryById",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.GetStoryById",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.GetStoryById(context)
			// assertions

		})
	}

}

// UpdateStory tests
func TestUpdateStory(t *testing.T) {
	storyController := StoryControllerImpl{}
	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusCreated,
			expectedResponseBody: []byte("pong"),
			testMessage:          "Happy path for StoryController.UpdateStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.UpdateStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.UpdateStory(context)
			// assertions

		})
	}
}

// deleteStory tests
func TestDeleteStory(t *testing.T) {
	storyController := StoryControllerImpl{}
	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNoContent,
			expectedResponseBody: []byte("pong"),
			testMessage:          "Happy path for StoryController.deleteStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.deleteStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.DeleteStory(context)
			// assertions

		})
	}
}

func TestGetStoriesByAuthor(t *testing.T) {
	// We can create the expected responses as an array of structs
	storyController := StoryControllerImpl{}
	testCases := []struct {
		writer               *httptest.ResponseRecorder
		expectedResponseCode int
		expectedResponseBody []byte
		testMessage          string
	}{
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte("pong"),
			testMessage:          "Happy path for StoryController.GetStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.GetStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.GetStoriesByAuthor(context)
			// assertions

		})
	}
}*/
