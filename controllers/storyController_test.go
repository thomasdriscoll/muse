package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Global variables
var storyId = 1 // 1 until I have something better

func TestGetStory(t *testing.T) {
	// We can create the expected responses as an array of structs
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
			testMessage:          "Happy path for StoryController.getStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.getStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			getRandomStory(context)
			// assertions
		})
	}
}

func TestCreateStory(t *testing.T) {
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
			testMessage:          "Happy path for StoryController.createStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.createStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story request not valid for StoryController.createStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			createStory(context)
			// assertions

		})
	}

}

func TestGetStoryById(t *testing.T) {
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
			testMessage:          "Happy path for StoryController.getStoryById",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.getStoryById",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			getStoryById(context)
			// assertions

		})
	}

}

// updateStory tests
func TestUpdateStory(t *testing.T) {
	testCases := []struct {
		writer               *httptest.ResponseRecorder
		expectedResponseCode int
		expectedResponseBody []byte
		testMessage          string
	}{
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusCreated,
			expectedResponseBody: []byte("pong"),
			testMessage:          "Happy path for StoryController.updateStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.updateStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			updateStory(context)
			// assertions

		})
	}
}

// deleteStory tests
func TestDeleteStory(t *testing.T) {
	testCases := []struct {
		writer               *httptest.ResponseRecorder
		expectedResponseCode int
		expectedResponseBody []byte
		testMessage          string
	}{
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
			deleteStory(context)
			// assertions

		})
	}
}

func TestGetStoriesByAuthor(t *testing.T) {
	// We can create the expected responses as an array of structs
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
			testMessage:          "Happy path for StoryController.getStory",
		},
		{
			writer:               httptest.NewRecorder(),
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.getStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			getStoriesByAuthor(context)
			// assertions

		})
	}
}
