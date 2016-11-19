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
	SearchFor(tag string) []Job
}

type BerlinStartupJobs struct {
	url string
}

type StackOverflow struct {
	url string
}

func (w BerlinStartupJobs) SearchFor(tag string) []Job {
	resp, err := http.Get(w.url + "/?s=" + tag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	//define matcher to find the job title
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent, "class") == "product-listing-h2"
		}
		return false
	}

	jobs := make([]Job, 0)
	articles := scrape.FindAll(root, matcher)
	for _, article := range articles {
		jobs = append(jobs, Job{Title: scrape.Text(article), Url: scrape.Attr(article, "href")})
	}

	return jobs
}

func (w StackOverflow) SearchFor(tag string) []Job {
	resp, err := http.Get(w.url + "/jobs?sort=p&q=" + tag + "&l=Berlin%2C+Germany&d=100&u=Km")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	//define matcher to find the job title
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil && scrape.Attr(n.Parent, "class") != "pagination" {
			return scrape.Attr(n, "class") == "job-link"
		}
		return false
	}

	jobs := make([]Job, 0)
	articles := scrape.FindAll(root, matcher)
	for _, article := range articles {
		jobs = append(jobs, Job{Title: scrape.Text(article), Url: w.url + scrape.Attr(article, "href")})
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
