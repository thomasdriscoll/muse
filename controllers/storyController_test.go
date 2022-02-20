package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
	"github.com/thomasdriscoll/muse/testhelper"
)

// Global variables
var storyId = 1 // 1 until I have something better
var storyController StoryController
var story = models.Story{
	Metadata: models.StoryMetadata{},
	Content:  getStoryContent(),
}

func getStoryContent() string {
	content, err := os.ReadFile(testhelper.GetTextFilePath())
	if err != nil {
		jsonContent, _ := json.Marshal(content)
		return jsonContent
	} else {
		error.New("whoopsie goof, you messed up good on the testdata")
	}
}

type TestCase struct {
	writer               *httptest.ResponseRecorder
	request              *http.Request
	expectedResponseCode int
	expectedResponseBody []byte
	testMessage          string
}

// *************************************************************************************************************
// 				TESTS
// *************************************************************************************************************

func TestGetStory(t *testing.T) {
	// Constants
	route := "/story"

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests & responses
	getRandomStoryRequest, _ := http.NewRequest(http.MethodGet, route, nil)
	okResponse, _ := json.Marshal(story)
	notFoundResponse, _ := json.Marshal(enums.ErrorStoryNotFound)
	dbErrResponse, _ := json.Marshal(enums.ErrorDBError)

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
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}
		})
	}
}

func TestCreateStoryFromURL(t *testing.T) {
	// Constants
	route := "/story/createFromURL"

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests
	storyFromURLRequestNoId := models.StoryFromURLRequest{
		Author:   "Hemingway, Ernest",
		AuthorId: "",
		UrlType:  "Gutenberg",
		Url:      "https://www.gutenberg.org/cache/epub/67138/pg67138.txt",
	}
	createRequest, _ := http.NewRequest(http.MethodPost, route, string.NewReader(storyFromURLRequestNoId))
	createResponse, _ := json.Marshal(story)
	invalidURLResponse, _ := json.Marshal(enums.ErrorInvalidURL)
	dbErrorResponse, _ := json.Marshal(enums.ErrorDBError)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              createRequest,
			expectedResponseCode: http.StatusCreated,
			expectedResponseBody: []byte(createResponse),
			testMessage:          "Happy path for StoryController.CreateStoryFromURL",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              createRequest,
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte(invalidURLResponse),
			testMessage:          "Test that InvalidURLs are rejected",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              createRequest,
			expectedResponseCode: http.StatusServiceUnavailable,
			expectedResponseBody: []byte(dbErrorResponse),
			testMessage:          "Database error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			context.Request = testCase.request
			storyController.CreateStoryFromURL(context)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}
		})
	}

}

func TestGetStoryById(t *testing.T) {
	// Constants
	route := "/story/storyId/"
	storyId := 1

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests & responses
	getStoryByIdRequest, _ := http.NewRequest(http.MethodGet, route+strconv.Itoa(storyId), nil)
	okResponse, _ := json.Marshal(story)
	notFoundResponse, _ := json.Marshal(enums.ErrorStoryNotFound)
	dbErrResponse, _ := json.Marshal(enums.ErrorDBError)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getStoryByIdRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(okResponse),
			testMessage:          "Happy path for StoryController.GetStoryById",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoryByIdRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Story not found for StoryController.GetStoryById",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoryByIdRequest,
			expectedResponseCode: http.StatusServiceUnavailable,
			expectedResponseBody: []byte(dbErrResponse),
			testMessage:          "Database error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.GetStoryById(context)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}
		})
	}

}

// deleteStory tests
func TestDeleteStory(t *testing.T) {
	// Constants
	route := "/story/storyId/"
	storyId := 1

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests & responses
	deleteStoryByIdRequest, _ := http.NewRequest(http.MethodDelete, route+strconv.Itoa(storyId), nil)
	notFoundResponse, _ := json.Marshal(enums.ErrorStoryNotFound)
	dbErrResponse, _ := json.Marshal(enums.ErrorDBError)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              deleteStoryByIdRequest,
			expectedResponseCode: http.StatusNoContent,
			expectedResponseBody: []byte(""),
			testMessage:          "Happy path for StoryController.deleteStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              deleteStoryByIdRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Story not found for StoryController.deleteStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              deleteStoryByIdRequest,
			expectedResponseCode: http.StatusServiceUnavailable,
			expectedResponseBody: []byte(dbErrResponse),
			testMessage:          "Database not available",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.DeleteStory(context)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}

		})
	}
}

func TestGetStoriesByAuthor(t *testing.T) {
	// Constants
	route := "/story/authors/"
	authorId := 1

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests & responses
	getStoriesByAuthorRequest, _ := http.NewRequest(http.MethodGet, route+strconv.Itoa(authorId), nil)
	multipleStoriesResponse, _ := json.Marshal([]models.Story{story})
	notFoundResponse, _ := json.Marshal(enums.ErrorAuthorNotFound)
	dbErrResponse, _ := json.Marshal(enums.ErrorDBError)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByAuthorRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(multipleStoriesResponse),
			testMessage:          "Happy path for StoryController.GetStoriesByAuthor",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByAuthorRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.GetStoriesByAuthor",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByAuthorRequest,
			expectedResponseCode: http.StatusServiceUnavailable,
			expectedResponseBody: []byte(dbErrResponse),
			testMessage:          "Database not available for StoryController.GetStoriesByAuthor",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.GetStoriesByAuthor(context)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}

		})
	}
}
func TestGetStoriesByTag(t *testing.T) {
	// Constants
	route := "/story/authors/"
	tag := "science-fiction"

	//Mocks
	storyController := StoryControllerImpl{}

	// Requests & responses
	getStoriesByAuthorRequest, _ := http.NewRequest(http.MethodGet, route+tag, nil)
	multipleStoriesResponse, _ := json.Marshal([]models.Story{story})
	notFoundResponse, _ := json.Marshal(enums.ErrorTagNotFound)
	dbErrResponse, _ := json.Marshal(enums.ErrorDBError)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByAuthorRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(multipleStoriesResponse),
			testMessage:          "Happy path for StoryController.GetStoriesByTag",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByAuthorRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte("you done goofed kid"),
			testMessage:          "Story not found for StoryController.GetStoriesByTag",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByAuthorRequest,
			expectedResponseCode: http.StatusServiceUnavailable,
			expectedResponseBody: []byte(dbErrResponse),
			testMessage:          "Database not available for StoryController.GetStoriesByTag",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			storyController.GetStoriesByAuthor(context)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Response body does not match expected response")
			}

		})
	}
}
