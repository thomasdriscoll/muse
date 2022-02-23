package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/controllers"
)

type App struct {
	Router *gin.Engine
}

// Router setup
func setupRouter() *gin.Engine {
	r := gin.Default()
	initializeRoutes(r)
	return r
}

// Add each controller here
func initializeRoutes(r *gin.Engine) {
	storyPrefix := r.Group("/story")
	storyController := controllers.StoryControllerImpl{}
	StoryRouteHandler(storyPrefix, storyController)

	userPrefix := r.Group("/user")
	userController := controllers.UserControllerImpl{}
	UserRouteHandler(userPrefix, userController)
}

func main() {
	app := App{
		Router: setupRouter(),
	}

	app.Router.Run()
}
