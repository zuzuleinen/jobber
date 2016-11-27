package sources

import (
	"golang.org/x/net/html"
)

type Source interface {
	Name() string
	QueryUrl(tag string) string
	Matcher() func(n *html.Node) bool
	Jobs(root *html.Node, tag string) []Job
}

//Get all available sources
func All() []Source {
	s := make([]Source, 0)
	s = append(s, BerlinStartupJobs{"http://berlinstartupjobs.com"})
	s = append(s, StackOverflow{"http://stackoverflow.com"})
	return s
}
