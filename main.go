package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

// All initialization logic here
func (a *App) initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Add each controller here
func (a *App) initializeRoutes() {
	StoryRouteHandler(a.Router)
	UserRouteHandler(a.Router)
}

// All run logic here
func (a *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	app := App{}
	app.initialize()

	app.run(":8080")
}
