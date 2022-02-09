package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

// Add each controller here
func (a *App) initializeRoutes() {
	storyPrefix := a.Router.Group("/story")
	a.Router.StoryRouteHandler(storyPrefix)

	userPrefix := a.Router.Group("/user")
	a.Router.UserRouteHandler(userPrefix)
}

func main() {
	app := App{
		Router: gin.Default(),
	}
	app.initializeRoutes()

	app.Router.Run()
}
