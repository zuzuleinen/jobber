package sources

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
)

type Job struct {
	Title string
	Url   string
}

type Source interface {
	QueryUrl(tag string) string
	Matcher() func(n *html.Node) bool
}

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

type StackOverflow struct {
	url string
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
func (w StackOverflow) QueryUrl(tag string) string {
	return w.url + "/jobs?sort=p&q=" + tag + "&l=Berlin%2C+Germany&d=100&u=Km"
}

//Search for a tag in a Source
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
	jobs := make([]Job, 0)
	articles := scrape.FindAll(root, s.Matcher())
	for _, article := range articles {
		jobs = append(jobs, Job{Title: scrape.Text(article), Url: scrape.Attr(article, "href")})
	}

	return jobs
}

//Get all available sources
func All() []Source {
	s := make([]Source, 0)
	s = append(s, BerlinStartupJobs{"http://berlinstartupjobs.com"})
	s = append(s, StackOverflow{"http://stackoverflow.com"})
	return s
}
