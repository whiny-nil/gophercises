package main

import (
	"html/template"
	"net/http"
	"os"
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

	rawTemplate, err := os.ReadFile("templates/story.tmpl.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("story").Parse(string(rawTemplate))
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, bh.storyJson[key])
}

func newStoryHandler(storyJson Story) *StoryHandler {
	sh := new(StoryHandler)
	sh.storyJson = storyJson

	return sh
}
