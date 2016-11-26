package sources

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type BerlinStartupJobs struct {
	url string
}

func (w BerlinStartupJobs) QueryUrl(tag string) string {
	return w.url + "/?s=" + tag
}
func (s BerlinStartupJobs) Matcher() func(n *html.Node) bool {
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent, "class") == "product-listing-h2"
		}
		return false
	}
	return matcher
}
func (s BerlinStartupJobs) Jobs(root *html.Node, tag string) []Job {
	jobs := make([]Job, 0)
	titles := scrape.FindAll(root, s.Matcher())
	for _, title := range titles {
		dateAdded := scrape.Text(title.Parent.NextSibling.NextSibling.NextSibling.NextSibling)
		jobs = append(jobs, Job{Title: scrape.Text(title), Url: scrape.Attr(title, "href"), Tag: tag, DateAdded:dateAdded})
	}

	return jobs
}

type StackOverflow struct {
	url string
}

func (w StackOverflow) QueryUrl(tag string) string {
	return w.url + "/jobs?sort=p&q=" + tag + "&l=Berlin%2C+Germany&d=100&u=Km"
}

func (s StackOverflow) Matcher() func(n *html.Node) bool {
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil && scrape.Attr(n.Parent, "class") != "pagination" {
			return scrape.Attr(n, "class") == "job-link"
		}
		return false
	}
	return matcher
}

func (s StackOverflow) Jobs(root *html.Node, tag string) []Job {
	jobs := make([]Job, 0)
	titles := scrape.FindAll(root, s.Matcher())
	for _, title := range titles {
		dateNode := title.
		Parent.
			Parent.
			NextSibling.
			NextSibling.
			NextSibling.
			NextSibling.
			NextSibling.
			NextSibling.
			NextSibling.
			NextSibling.
			NextSibling
		jobs = append(jobs, Job{Title: scrape.Text(title), Url: s.url + scrape.Attr(title, "href"), Tag: tag, DateAdded:scrape.Text(dateNode)})
	}

	return jobs
}
