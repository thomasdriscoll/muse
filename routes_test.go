package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

type TestCase struct {
	writer               *httptest.ResponseRecorder
	request              *http.Request
	expectedResponseCode int
	expectedResponseBody []byte
	testMessage          string
}

/*
This test validates each of my routes in StoryController and makes sure that a Gin engine would actually be able to reach them
Mocks the handlers, those are tested in the controllers themselves
*/
func TestStoryRouteHandler(t *testing.T) {
	// constants
	routePrefix := "/story"
	storyId := 1
	authorId := 2
	tagId := "science"

	// requests
	// NOTE: Two return values meant they couldn't be inline in the testCases slice
	getRandomStoryRequest, _ := http.NewRequest(http.MethodGet, routePrefix, nil)
	createStoryFromURLRequest, _ := http.NewRequest(http.MethodPost, routePrefix+"/createFromURL", nil) // add CreateFromURL request here
	getStoryByIdRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/storyId/"+strconv.Itoa(storyId), nil)
	deleteStoryRequest, _ := http.NewRequest(http.MethodDelete, routePrefix+"/storyId/"+strconv.Itoa(storyId), nil)
	getStoryByAuthorRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/authors/"+strconv.Itoa(authorId), nil)
	getStoriesByTagRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/tag/"+tagId, nil)
	junkRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/junk", nil)

	// mocks
	mockStoryController := MockStoryController{}
	r := gin.Default()
	routeGroup := r.Group(routePrefix)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Actual test
	StoryRouteHandler(routeGroup, mockStoryController)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              getRandomStoryRequest,
			expectedResponseCode: 200,
			expectedResponseBody: []byte("\"Get Random Story works\""),
			testMessage:          "GET /story route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              createStoryFromURLRequest,
			expectedResponseCode: 201,
			expectedResponseBody: []byte("\"Create Story From URL works\""),
			testMessage:          "POST /story/createFromURL route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoryByIdRequest,
			expectedResponseCode: 200,
			expectedResponseBody: []byte("\"Get Story By Id works\""),
			testMessage:          "GET /story/:id route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              deleteStoryRequest,
			expectedResponseCode: 204,
			expectedResponseBody: []byte(""),
			testMessage:          "DELETE /story/:id route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoryByAuthorRequest,
			expectedResponseCode: 200,
			expectedResponseBody: []byte("\"Get Stories By Author works\""),
			testMessage:          "GET /story/authors/:authorId route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getStoriesByTagRequest,
			expectedResponseCode: 200,
			expectedResponseBody: []byte("\"Get Stories By Tag works\""),
			testMessage:          "GET /story/tag/:tagId route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              junkRequest,
			expectedResponseCode: 404,
			expectedResponseBody: []byte("404 page not found"),
			testMessage:          "404 on /story/junk to make sure no false positives",
		},
	}

	// test cases
	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			r.ServeHTTP(testCase.writer, testCase.request)
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Status Code didn't match:\n\t%q\n\t%q", testCase.expectedResponseCode, testCase.writer.Code)
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Body didn't match:\n\t%q\n\t%q", string(testCase.expectedResponseBody), testCase.writer.Body.String())
			}
		})
	}
}

func TestUserRouteHandler(t *testing.T) {
	// constants
	routePrefix := "/user"
	userId := 1

	// requests
	createUserRequest, _ := http.NewRequest(http.MethodPost, routePrefix, nil) // Add User Context here
	getUserRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/userId/"+strconv.Itoa(userId), nil)
	getSavedStoriesByUserRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/userId/savedStories"+strconv.Itoa(userId), nil)
	junkRequest, _ := http.NewRequest(http.MethodGet, routePrefix+"/junk", nil)

	// mocks
	mockUserController := MockUserController{}
	r := gin.Default()
	routeGroup := r.Group(routePrefix)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Actual test
	UserRouteHandler(routeGroup, mockUserController)

	testCases := []TestCase{
		{
			writer:               httptest.NewRecorder(),
			request:              createUserRequest,
			expectedResponseCode: 201,
			expectedResponseBody: []byte("\"works\""),
			testMessage:          "POST /user route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getUserRequest,
			expectedResponseCode: 200,
			expectedResponseBody: []byte("\"works\""),
			testMessage:          "GET /user/:userId route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              getSavedStoriesByUserRequest,
			expectedResponseCode: 200,
			expectedResponseBody: []byte("\"works\""),
			testMessage:          "GET /user/:userId/savedStories route",
		},
		{
			writer:               httptest.NewRecorder(),
			request:              junkRequest,
			expectedResponseCode: 404,
			expectedResponseBody: []byte("404 page not found"),
			testMessage:          "GET /user/junk route",
		},
	}

	// Run test suite
	for _, testCase := range testCases {
		t.Run(testCase.testMessage, func(t *testing.T) {
			r.ServeHTTP(testCase.writer, testCase.request)
			if testCase.expectedResponseCode != testCase.writer.Code {
				t.Errorf("Status Code didn't match:\n\t%q\n\t%q", testCase.expectedResponseCode, testCase.writer.Code)
			}
			if !bytes.Equal(testCase.expectedResponseBody, testCase.writer.Body.Bytes()) {
				t.Errorf("Body didn't match:\n\t%q\n\t%q", string(testCase.expectedResponseBody), testCase.writer.Body.String())
			}
		})
	}

}

func TestBadRoutes(t *testing.T) {

}

// All mocking structs and related functions
// Note: interfaces for these structs exist in respective code base with impl

// StoryController stubs
type MockStoryController struct{}

func (msc MockStoryController) GetRandomStory(c *gin.Context) {
	c.JSON(200, "Get Random Story works")

}

func (msc MockStoryController) CreateStoryFromURL(c *gin.Context) {
	c.JSON(201, "Create Story From URL works")

}

func (msc MockStoryController) GetStoryById(c *gin.Context) {
	c.JSON(200, "Get Story By Id works")

}

func (msc MockStoryController) DeleteStory(c *gin.Context) {
	c.JSON(204, "")

}

func (msc MockStoryController) GetStoriesByAuthor(c *gin.Context) {
	c.JSON(200, "Get Stories By Author works")

}

func (msc MockStoryController) GetStoriesByTag(c *gin.Context) {
	c.JSON(200, "Get Stories By Tag works")
}

// UserController stubs
type MockUserController struct{}

func (msc MockUserController) CreateUser(c *gin.Context) {
	c.JSON(201, "works")
}

func (muc MockUserController) GetUser(c *gin.Context) {
	c.JSON(200, "works")
}

func (muc MockUserController) GetSavedStoriesByUser(c *gin.Context) {
	c.JSON(200, "works")
}
