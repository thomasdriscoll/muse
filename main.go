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

func (a *App) initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
}

func (a *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	app := App{}
	app.initialize()

	app.run(":8080")
}
