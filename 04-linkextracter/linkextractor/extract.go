package linkextractor

type Link struct {
	Href string
	Text string
}

func Extract(html string) []Link {
	retVal := []Link{}

	retVal = append(retVal, Link{"/link", "Link text"})

	return retVal
}
