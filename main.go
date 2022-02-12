package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/thomasdriscoll/muse/controllers"
)

type App struct {
	DB     *sql.DB
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
	controllers.StoryRouteHandler(storyPrefix)

	userPrefix := r.Group("/user")
	controllers.UserRouteHandler(userPrefix)
}

func main() {
	app := App{
		Router: setupRouter(),
		DB:     nil,
	}

	app.Router.Run()
}
