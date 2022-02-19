package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/controllers"
)

// Potential tech debt -- abstract every controller to share interface, too complicated for now (for me)
type StoryController interface {
	GetRandomStory(c *gin.Context)
	CreateStoryFromURL(c *gin.Context)
	CreateStoryFromFile(c *gin.Context)
	GetStoryById(c *gin.Context)
	UpdateStory(c *gin.Context)
	DeleteStory(c *gin.Context)
	GetStoriesByAuthor(c *gin.Context)
	GetStoriesByTag(c *gin.Context)
}

type StoryControllerImpl struct{}

func StoryRouteHandler(story *gin.RouterGroup, storyController controllers.StoryController) {
	story.GET("", storyController.GetRandomStory)
	story.POST("/createFromURL", storyController.CreateStoryFromURL)
	story.POST("/createFromFile", storyController.CreateStoryFromFile)
	story.GET("/storyId/:id", storyController.GetStoryById)
	story.PUT("/storyId/:id", storyController.UpdateStory)
	story.DELETE("/storyId/:id", storyController.DeleteStory)
	story.GET("/authors/:authorId", storyController.GetStoriesByAuthor)
	story.GET("/tag/:tag", storyController.GetStoriesByTag)
}

func (sc StoryControllerImpl) GetRandomStory(c *gin.Context) {
	c.JSON(200, "pong")
}

func (sc StoryControllerImpl) CreateStoryFromURL(c *gin.Context) {
}

func (sc StoryControllerImpl) CreateStoryFromFile(c *gin.Context) {
}

func (sc StoryControllerImpl) GetStoryById(c *gin.Context) {
	c.JSON(200, "testing")
}

func (sc StoryControllerImpl) UpdateStory(c *gin.Context) {
}

func (sc StoryControllerImpl) DeleteStory(c *gin.Context) {
}

func (sc StoryControllerImpl) GetStoriesByAuthor(c *gin.Context) {

}

func (sc StoryControllerImpl) GetStoriesByTag(c *gin.Context) {
}
