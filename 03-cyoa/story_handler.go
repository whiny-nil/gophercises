package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type StoryHandler struct {
	storyJson Story
}

func (bh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Trim(r.URL.Path, "/")
	if key == "" {
		key = "intro"
	}
	title := bh.storyJson[key].Title
	log.Printf("%s", title)

	fmt.Fprint(w, title)
}

func newStoryHandler(storyJson Story) *StoryHandler {
	sh := new(StoryHandler)
	sh.storyJson = storyJson

	return sh
}
