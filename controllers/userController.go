package main

import "github.com/gorilla/mux"

type UserController struct {
}

func UserRouteHandler(r *mux.Router) {
	subrouter := r.PathPrefix("/user").Subrouter()
	subrouter.HandleFunc("/", getUser).Methods("POST")
	subrouter.HandleFunc("/{userId}", createUser).Methods("GET")
	subrouter.HandleFunc("/{userId}/savedStories", getSavedStoriesByUser).Methods("GET")
}
