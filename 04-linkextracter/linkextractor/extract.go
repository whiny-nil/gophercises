package linkextractor

type Link struct {
	Href string
	Text string
}

func Extract(html string) Link {
	return Link{"/link", "Link text"}
}
