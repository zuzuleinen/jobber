package sources

import (
	"golang.org/x/net/html"
	"net/http"
)

type Job struct {
	Title string
	Url   string
	Tag   string
	DateAdded string
}

//Search for jobs related to tag in a Source
func SearchFor(tag string, s Source) []Job {
	resp, err := http.Get(s.QueryUrl(tag))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	return s.Jobs(root, tag)
}
