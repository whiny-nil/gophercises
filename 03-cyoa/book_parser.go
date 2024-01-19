package main

import (
	"encoding/json"
	"os"
)

type Title string
type Paragraphs []string
type Option struct {
	Text string
	Arc  string
}
type Options []Option
type Page struct {
	Title      Title
	Paragraphs Paragraphs
	Options    Options
}
type Story map[string]Page

func parseBook(fileName *string) ([]byte, error) {
	storyData, err := os.ReadFile(*storyFile)
	if err != nil {
		return nil, err
	}

	var storyJson Story
	err = json.Unmarshal(storyData, &storyJson)
	if err != nil {
		return nil, err
	}

	return storyData, nil
}
