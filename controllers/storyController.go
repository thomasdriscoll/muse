package main

import (
	"github.com/gin-gonic/gin"
)

type StoryController struct {
}

func (*gin.Engine) StoryRouteHandler(rg *gin.RouterGroup) {
	story := rg.Group("/story")

	story.GET("/", getRandomStory)
	story.POST("/", createStory)
	story.GET("/:id", getStoryById)
	story.PUT("/:id", updateStory)
	story.DELETE("/:id", deleteStory)
	story.GET("/authors/:authorId", getStoriesByAuthor)
	story.GET("/tag/:tag", getStoryById)
}

func getRandomStory(c *gin.Context) {
	TODO()
}

func createStory(c *gin.Context) {
	TODO()
}

func getStoryByIdStory(c *gin.Context) {
	TODO()
}

func updateStory(c *gin.Context) {
	TODO()
}

func deleteStory(c *gin.Context) {
	TODO()
}

func getStoriesByAuthor(c *gin.Context) {
	TODO()
}

func getStoryById(c *gin.Context) {
	TODO()
}
