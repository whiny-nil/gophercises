package main

import (
	"fmt"
	"net/http"
)

type BookHandler struct{}

func (bh BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func newBookHandler() *BookHandler {
	return new(BookHandler)
}
