package sources

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strconv"
	"strings"
	"time"
)

type Source interface {
	Name() string
	QueryUrl(tag string) string
	Matcher() func(n *html.Node) bool
	Jobs(root *html.Node, tag string) []Job
}

type BerlinStartupJobs struct {
	url string
}

func (w BerlinStartupJobs) Name() string {
	return "berlinstartupjobs"
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
		t, err := time.Parse("January 2, 2006", dateAdded)

		if err != nil {
			panic(err)
		}

		jobs = append(
			jobs,
			Job{
				Title:     scrape.Text(title),
				Url:       scrape.Attr(title, "href"),
				Tag:       tag,
				DateAdded: t.Format(time.RFC822),
			},
		)
	}

	return jobs
}

type StackOverflow struct {
	url string
}

func (w StackOverflow) Name() string {
	return "stackoverflow"
}
func (w StackOverflow) QueryUrl(tag string) string {
	return w.url + "/jobs?sort=p&q=" + tag + "&l=Berlin%2C+Germany&d=100&u=Km"
}
func (s StackOverflow) Matcher() func(n *html.Node) bool {
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil && scrape.Attr(n.Parent, "class") != "pagination" {
			if scrape.Attr(n, "class") == "job-link" {
				//remove featured jobs since they are not sorted by date
				if strings.Contains(scrape.Attr(n.Parent.Parent.Parent, "class"), "highlighted") {
					return false
				}
				return true
			}
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

		t := time.Now()
		duration, err := time.ParseDuration(s.FormatDate(scrape.Text(dateNode)))

		if err != nil {
			panic(err)
		}

		jobs = append(
			jobs,
			Job{
				Title:     scrape.Text(title),
				Url:       s.url + scrape.Attr(title, "href"),
				Tag:       tag,
				DateAdded: t.Add(duration).Format(time.RFC822),
			},
		)
	}

	return jobs
}

func (s StackOverflow) FormatDate(dateAdded string) string {
	if strings.Contains(dateAdded, "hours") {
		dateAdded = strings.Replace(dateAdded, "hours", "h", -1)
	}
	if strings.Contains(dateAdded, "hour") {
		dateAdded = strings.Replace(dateAdded, "hour", "h", -1)
	}
	if strings.Contains(dateAdded, "yesterday") {
		dateAdded = strings.Replace(dateAdded, "yesterday", "24h", -1)
	}
	if strings.Contains(dateAdded, "days") || strings.Contains(dateAdded, "week") {
		elements := strings.Split(dateAdded, " ")

		counter, _ := strconv.Atoi(elements[0])
		unit := elements[1]

		hours := 0
		if strings.Contains(unit, "day") {
			hours = counter * 24
		}
		if strings.Contains(unit, "week") {
			hours = counter * 7 * 24
		}
		dateAdded = strconv.Itoa(hours) + "h"
	}
	dateAdded = strings.Replace(dateAdded, "ago", "", -1)
	dateAdded = strings.Replace(dateAdded, " ", "", -1)
	dateAdded = strings.Replace(dateAdded, "<", "", -1)
	durationString := "-" + dateAdded

	return durationString
}
