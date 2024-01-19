package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var storyFile = flag.String("f", "stories/gopher.json", "a story file to use")

func main() {
	flag.Parse()

	storyJson, err := parseStory(storyFile)
	if err != nil {
		log.Fatal(err)
	}

	storyHandler := newStoryHandler(storyJson)
	mux := http.NewServeMux()
	mux.Handle("/", storyHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}
