package controllers

import (
	"github.com/gin-gonic/gin"
)

type StoryController interface {
}

func StoryRouteHandler(story *gin.RouterGroup) {
	story.GET("/", GetRandomStory)
	story.POST("/", CreateStory)
	story.GET("/:id", GetStoryById)
	story.PUT("/:id", UpdateStory)
	story.DELETE("/:id", DeleteStory)
	story.GET("/authors/:authorId", GetStoriesByAuthor)
	story.GET("/tag/:tag", GetStoryById)
}

func (sc *StoryController) GetRandomStory(c *gin.Context) {
	c.JSON(200, "pong")
}

func (sc *StoryController) CreateStory(c *gin.Context) {
}

func (sc *StoryController) GetStoryById(c *gin.Context) {

}

func (sc *StoryController) UpdateStory(c *gin.Context) {
}

func (sc *StoryController) DeleteStory(c *gin.Context) {
}

func (sc *StoryController) GetStoriesByAuthor(c *gin.Context) {
}
