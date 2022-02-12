package main

import "github.com/gin-gonic/gin"

func StoryRouteHandler(story *gin.RouterGroup) {
	storyController := StoryController{}
	story.GET("/", storyController.GetRandomStory)
	story.POST("/", storyController.CreateStory)
	story.GET("/:id", storyController.GetStoryById)
	story.PUT("/:id", storyController.UpdateStory)
	story.DELETE("/:id", storyController.DeleteStory)
	story.GET("/authors/:authorId", storyController.GetStoriesByAuthor)
	story.GET("/tag/:tag", storyController.GetStoryById)
}

func UserRouteHandler(user *gin.RouterGroup) {
	userController := UserController{}
	user.GET("/", userController.GetUser)
	user.POST("/:userId", userController.CreateUser)
	user.GET("/:userId/savedStories", userController.GetSavedStoriesByUser)
}
