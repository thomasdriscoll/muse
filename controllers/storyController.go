package controllers

import (
	"github.com/gin-gonic/gin"
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
