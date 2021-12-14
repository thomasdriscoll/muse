package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

}

func main() {
	a := App{}
	a.Intialize()
	a.Run(":8010")
}
