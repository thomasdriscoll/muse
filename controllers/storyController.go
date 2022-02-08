package main

import "github.com/gorilla/mux"

type StoryController struct {
}

func StoryRouteHandler(r *mux.Router) {
	subrouter := r.PathPrefix("/story").Subrouter()
	subrouter.HandleFunc("/", getRandomStory).Methods("GET")
	subrouter.HandleFunc("/", createStory).Methods("POST")
	subrouter.HandleFunc("/{id}", getStoryById).Methods("GET")
	subrouter.HandleFunc("/{id}", updateStory).Methods("PUT")
	subrouter.HandleFunc("/{id}", deleteStory).Methods("DELETE")
	subrouter.HandleFunc("/authors/{authorId}", getStoriesByAuthor).Methods("GET")
	subrouter.HandleFunc("/tag/{tag}", getStoryById).Methods("GET")
}
