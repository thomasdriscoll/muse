package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/story", nil)
	w := httptest.NewRecorder()
	getStory(w, req)
	result := w.Result()
	defer result.Body.Close()
}
