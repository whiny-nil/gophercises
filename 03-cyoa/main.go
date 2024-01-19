package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type helloWorldHandler struct{}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func NewHelloWorldHandler() *helloWorldHandler {
	return new(helloWorldHandler)
}

type Title string
type Story []string
type Option struct {
	Text string
	Arc  string
}
type Options []Option
type Page struct {
	Title   Title
	Story   Story
	Options Options
}
type Book map[string]Page

var storyFile = flag.String("f", "stories/gopher.json", "a story file to use")

func main() {
	flag.Parse()

	content, err := os.ReadFile(*storyFile)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var book Book
	err = json.Unmarshal(content, &book)
	if err != nil {
		log.Fatal("Error unmarshalling data: ", err)
	}

	log.Printf("intro: \n%s", book["intro"])

	helloHandler := NewHelloWorldHandler()
	mux := http.NewServeMux()
	mux.Handle("/", helloHandler)

	// 	// Build the MapHandler using the mux as the fallback
	// 	pathsToUrls := map[string]string{
	// 		"/urlshort-godoc": "https://github.com/gophercises/urlshort",
	// 		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// 	}
	// 	mapHandler := urlhandler.MapHandler(pathsToUrls, mux)

	// 	// Build the YAMLHandler using the mapHandler as the fallback
	// 	yaml := `
	// - path: "/urlshort"
	//   url: "https://github.com/gophercises/urlshort"
	// - path: "/urlshort-final"
	//   url: "https://github.com/gophercises/urlshort/tree/solution"
	// `
	// 	yamlHandler, err := urlhandler.YAMLHandler([]byte(yaml), mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}
