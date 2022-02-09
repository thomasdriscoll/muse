package main

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (*gin.Engine) UserRouteHandler(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/", getUser)
	user.POST("/{userId}", createUser)
	user.GET("/{userId}/savedStories", getSavedStoriesByUser)
}

func getUser(c *gin.Context) {
	TODO()
}

func createUser(c *gin.Context) {
	TODO()
}

func getSavedStoriesByUser(c *gin.Context) {
	TODO()
}
