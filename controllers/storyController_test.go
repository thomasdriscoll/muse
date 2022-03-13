package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
var storyId uint64 = 1 // 1 until I have something better
var story = models.Story{
	StoryMetadata: models.StoryMetadata{},
	Content:       getStoryContent(),
}

func getStoryContent() []byte {
	path, pathErr := testhelper.GetTextFilePath()
	if pathErr != nil {
		panic(pathErr.Error())
	}
	content, err := os.ReadFile(path)
	if err != nil {
		jsonContent, _ := json.Marshal(content)
		return jsonContent
	} else {
		panic(errors.New("whoopsie goof, you messed up good on the testdata"))
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
	storyController := StoryControllerImpl{
		StorySvc: &MockStoryService{},
	}
	engine := gin.New()
	engine.GET(route, storyController.GetRandomStory)

	// Requests & responses
	getRandomStoryRequest, _ := http.NewRequest(http.MethodGet, route, nil)
	okResponse, _ := json.Marshal(story)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getRandomStoryRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(okResponse),
			testMessage:          "Happy path for StoryController.GetStory",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			engine.ServeHTTP(testCase.writer, testCase.request)
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
	storyController := StoryControllerImpl{
		StorySvc: &MockStoryService{},
	}
	engine := gin.New()
	engine.POST(route, storyController.CreateStoryFromURL)

	// Requests
	badRequestBody := "junk"
	jsonBadRequest, _ := json.Marshal(badRequestBody)
	badRequest, _ := http.NewRequest(http.MethodPost, route, bytes.NewReader(jsonBadRequest))

	gutenbergRequestBody := models.StoryFromURLRequest{
		Tags:    []string{"realisticFiction"},
		Title:   "The Sun Also Rises",
		UrlType: enums.Gutenberg,
		Url:     "https://www.gutenberg.org/cache/epub/67138/pg67138.txt",
	}
	jsonGutenbergRequest, _ := json.Marshal(gutenbergRequestBody)
	gutenbergPostRequest, _ := http.NewRequest(http.MethodPost, route, bytes.NewReader(jsonGutenbergRequest))

	notFoundGutenbergRequestBody := models.StoryFromURLRequest{
		Tags:    []string{"realisticFiction"},
		Title:   "junk",
		UrlType: enums.Gutenberg,
		Url:     "https://www.gutenberg.org/cache/epub/67138/pg67138.txt",
	}
	jsonNotFoundGutenbergRequest, _ := json.Marshal(notFoundGutenbergRequestBody)
	notFoundGutenbergPostRequest, _ := http.NewRequest(http.MethodPost, route, bytes.NewReader(jsonNotFoundGutenbergRequest))

	wikipediaRequestBody := models.StoryFromURLRequest{
		Tags:    []string{"superhero"},
		Title:   "Superman",
		UrlType: enums.Wikipedia,
		Url:     "https://en.wikipedia.org/wiki/Superman",
	}
	jsonWikipediaRequest, _ := json.Marshal(wikipediaRequestBody)
	wikipediaPostRequest, _ := http.NewRequest(http.MethodPost, route, bytes.NewReader(jsonWikipediaRequest))

	notFoundWikipediaRequestBody := models.StoryFromURLRequest{
		Tags:    []string{"superhero"},
		Title:   "junk",
		UrlType: enums.Wikipedia,
		Url:     "https://en.wikipedia.org/wiki/Superman",
	}
	jsonNotFoundWikipediaRequest, _ := json.Marshal(notFoundWikipediaRequestBody)
	notFoundWikipediaPostRequest, _ := http.NewRequest(http.MethodPost, route, bytes.NewReader(jsonNotFoundWikipediaRequest))

	createResponse, _ := json.Marshal(story)
	badRequestResponse, _ := json.Marshal(enums.ErrorInvalidStoryRequest)
	notFoundURLResponse, _ := json.Marshal(enums.ErrorURLNotFound)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              badRequest,
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte(badRequestResponse),
			testMessage:          "Bad request StoryController.CreateStoryFromURL",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              gutenbergPostRequest,
			expectedResponseCode: http.StatusCreated,
			expectedResponseBody: []byte(createResponse),
			testMessage:          "Happy path for Gutenberg request StoryController.CreateStoryFromURL",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              notFoundGutenbergPostRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundURLResponse),
			testMessage:          "Test that InvalidURLs for Gutenberg are rejected",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              wikipediaPostRequest,
			expectedResponseCode: http.StatusCreated,
			expectedResponseBody: []byte(createResponse),
			testMessage:          "Happy path for Wikipedia request for StoryController.CreateStoryFromURL",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              notFoundWikipediaPostRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundURLResponse),
			testMessage:          "Test that InvalidURLs for Wikipedia are rejected",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			context, _ := gin.CreateTestContext(testCase.writer)
			context.Request = testCase.request
			storyController.CreateStoryFromURL(context)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				fmt.Println(testCase.expectedResponseCode)
				fmt.Println(testCase.writer.Code)
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				fmt.Println(string(testCase.expectedResponseBody))
				fmt.Println(testCase.writer.Body)
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
	storyController := StoryControllerImpl{
		StorySvc: &MockStoryService{},
	}
	engine := gin.New()
	engine.GET(route+":id", storyController.GetStoryById)

	// Requests & responses
	getStoryByIdRequest, _ := http.NewRequest(http.MethodGet, route+strconv.Itoa(storyId), nil)
	notFoundGetStoryByIdRequest, _ := http.NewRequest(http.MethodGet, route+strconv.Itoa(0), nil)
	invalidIdGetStoryByIdRequest, _ := http.NewRequest(http.MethodGet, route+"invalidId", nil)

	okResponse, _ := json.Marshal(story)
	notFoundResponse, _ := json.Marshal(enums.ErrorStoryNotFound)
	invalidIdResponse, _ := json.Marshal(enums.ErrorInvalidStoryId)

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
			request:              notFoundGetStoryByIdRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Story not found for StoryController.GetStoryById",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              invalidIdGetStoryByIdRequest,
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte(invalidIdResponse),
			testMessage:          "Invalid ID for StoryController.GetStoryById",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			engine.ServeHTTP(testCase.writer, testCase.request)
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
	storyController := StoryControllerImpl{
		StorySvc: &MockStoryService{},
	}
	engine := gin.New()
	engine.DELETE(route+":id", storyController.DeleteStory)

	// Requests & responses
	deleteStoryByIdRequest, _ := http.NewRequest(http.MethodDelete, route+strconv.Itoa(storyId), nil)
	notFoundDeleteRequest, _ := http.NewRequest(http.MethodDelete, route+strconv.Itoa(0), nil)
	invalidDeleteRequest, _ := http.NewRequest(http.MethodDelete, route+"junk", nil)

	notFoundResponse, _ := json.Marshal(enums.ErrorStoryNotFound)
	invalidIdResponse, _ := json.Marshal(enums.ErrorInvalidStoryId)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              deleteStoryByIdRequest,
			expectedResponseCode: http.StatusNoContent,
			expectedResponseBody: nil,
			testMessage:          "Happy path for StoryController.deleteStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              notFoundDeleteRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Story not found for StoryController.deleteStory",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              invalidDeleteRequest,
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte(invalidIdResponse),
			testMessage:          "Database not available",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			engine.ServeHTTP(testCase.writer, testCase.request)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				fmt.Println(string(testCase.expectedResponseBody))
				fmt.Println(testCase.writer.Body)
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
	storyController := StoryControllerImpl{
		StorySvc: &MockStoryService{},
	}
	engine := gin.New()
	engine.GET(route+":authorId", storyController.GetStoriesByAuthor)

	// Requests & responses
	getStoriesByAuthorRequest, _ := http.NewRequest(http.MethodGet, route+strconv.Itoa(authorId), nil)
	notFoundGetStoriesByAuthorIdRequest, _ := http.NewRequest(http.MethodGet, route+strconv.Itoa(0), nil)
	invalidIdGetStoriesByAuthorIdRequest, _ := http.NewRequest(http.MethodGet, route+"invalidId", nil)

	multipleStoriesResponse, _ := json.Marshal([]models.Story{story})
	notFoundResponse, _ := json.Marshal(enums.ErrorAuthorNotFound)
	invalidIdResponse, _ := json.Marshal(enums.ErrorInvalidStoryId)

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
			request:              notFoundGetStoriesByAuthorIdRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Story not found for StoryController.GetStoriesByAuthor",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              invalidIdGetStoriesByAuthorIdRequest,
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: []byte(invalidIdResponse),
			testMessage:          "Database not available for StoryController.GetStoriesByAuthor",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			engine.ServeHTTP(testCase.writer, testCase.request)
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
	route := "/story/tag/"
	tag := "scienceFiction"

	//Mocks
	storyController := StoryControllerImpl{
		StorySvc: &MockStoryService{},
	}
	engine := gin.New()
	engine.GET(route+":tag", storyController.GetStoriesByTag)

	// Requests & responses
	getStoriesByTagRequest, _ := http.NewRequest(http.MethodGet, route+tag, nil)
	notFoundGetStoriesByTagRequest, _ := http.NewRequest(http.MethodGet, route+"emptyTag", nil)

	multipleStoriesResponse, _ := json.Marshal([]models.Story{story})
	notFoundResponse, _ := json.Marshal(enums.ErrorTagNotFound)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByTagRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(multipleStoriesResponse),
			testMessage:          "Happy path for StoryController.GetStoriesByTag",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              notFoundGetStoriesByTagRequest,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: []byte(notFoundResponse),
			testMessage:          "Tag not found for StoryController.GetStoriesByTag",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			engine.ServeHTTP(testCase.writer, testCase.request)
			// assertions
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Response Code does not match expected status code")
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				fmt.Println(string(testCase.expectedResponseBody))
				fmt.Println(testCase.writer.Body)
				t.Errorf("Response body does not match expected response")
			}

		})
	}
}

// StorySvc stubs
type MockStoryService struct{}

func (r *MockStoryService) GetStoryById(ID uint64) (*models.Story, error) {
	if ID == storyId {
		return &story, nil
	} else {
		return nil, errors.New(enums.ErrorStoryNotFound)
	}
}

func (r *MockStoryService) CreateStory(storyRequest *models.StoryFromURLRequest) (*models.Story, error) {
	if storyRequest.Title == "junk" {
		return nil, errors.New(enums.ErrorURLNotFound)
	}
	return &story, nil
}

func (r *MockStoryService) GetStoryByRandom() (*models.Story, error) {
	return &story, nil
}

func (r *MockStoryService) DeleteById(ID uint64) error {
	if ID != storyId {
		return errors.New(enums.ErrorStoryNotFound)
	}
	return nil
}

func (r *MockStoryService) GetStoriesByTag(tag string) (*[]models.Story, error) {
	if tag == "scienceFiction" {
		stories := []models.Story{story}
		return &stories, nil
	}
	return nil, errors.New(enums.ErrorTagNotFound)
}

func (r *MockStoryService) GetStoriesByAuthorId(authorId uint64) (*[]models.Story, error) {
	if authorId == 1 {
		stories := []models.Story{story}
		return &stories, nil
	}
	return nil, errors.New(enums.ErrorAuthorNotFound)
}
