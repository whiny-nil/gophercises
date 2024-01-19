package main

import (
	"fmt"
	"net/http"
)

type helloWorldHandler struct{}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func NewHelloWorldHandler() *helloWorldHandler {
	return new(helloWorldHandler)
}

func main() {
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
