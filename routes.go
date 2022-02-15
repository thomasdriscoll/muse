package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/controllers"
)

func StoryRouteHandler(story *gin.RouterGroup, storyController controllers.StoryController) {
	story.GET("", storyController.GetRandomStory)
	story.POST("", storyController.CreateStory)
	story.GET("/storyId/:id", storyController.GetStoryById)
	story.PUT("/storyId/:id", storyController.UpdateStory)
	story.DELETE("/storyId/:id", storyController.DeleteStory)
	story.GET("/authors/:authorId", storyController.GetStoriesByAuthor)
	story.GET("/tag/:tag", storyController.GetStoriesByTag)
}

func UserRouteHandler(user *gin.RouterGroup, userController controllers.UserController) {
	user.POST("", userController.CreateUser)
	user.GET("/userId/:id", userController.GetUser)
	user.GET("/userId/savedStories/:id", userController.GetSavedStoriesByUser)
}
