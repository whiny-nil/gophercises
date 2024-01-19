package main

import (
	"fmt"
	"net/http"
)

type StoryHandler struct{}

func (bh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello world!")
}

func newStoryHandler() *StoryHandler {
	return new(StoryHandler)
}
