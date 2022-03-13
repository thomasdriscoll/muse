package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/thomasdriscoll/muse/controllers"
	"github.com/thomasdriscoll/muse/repositories"
	"github.com/thomasdriscoll/muse/services"
)

type App struct {
	Router *gin.Engine
}

// Router setup
func setup(db *pgx.Conn) *gin.Engine {
	// Create generics
	r := gin.Default()

	// Create services
	storyScrapper := services.StoryScrapperImpl{}

	// Create repositories
	storyRepo := repositories.NewStoryRepo(db)

	// Create controllers
	storyController := controllers.StoryControllerImpl{
		StoryRepo:     &storyRepo,
		StoryScrapper: &storyScrapper,
	}
	userController := controllers.UserControllerImpl{}

	storyPrefix := r.Group("/story")
	StoryRouteHandler(storyPrefix, storyController)
	userPrefix := r.Group("/user")
	UserRouteHandler(userPrefix, userController)
	return r
}

func main() {
	db := repositories.ConnectPostgreSQLDB()
	app := App{
		Router: setup(db),
	}

	app.Router.Run()
}
