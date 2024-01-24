package linkextracter_test

import (
	"testing"

	"linkextractor/linkextractor"
)

func TestSanity(t *testing.T) {
	if true != true {
		t.Error("true isn't equal to true!")
	}
}

func TestExtract(t *testing.T) {
	want := linkextractor.Link{Href: "/link", Text: "Link text"}

	got := linkextractor.Extract("<a href=\"/link\">Link text</a>")

	if want.Href != got.Href {
		t.Errorf("want %s, got %s", want.Href, got.Href)
	}

	if want.Text != got.Text {
		t.Errorf("want %s, got %s", want.Text, got.Text)
	}
}
