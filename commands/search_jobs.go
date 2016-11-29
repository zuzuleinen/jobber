package commands

import (
	"database/sql"
	"fmt"
	"github.com/zuzuleinen/jobber/database"
	"github.com/zuzuleinen/jobber/email"
	"github.com/zuzuleinen/jobber/sources"
	"time"
	"sort"
)

type ByDateDesc []sources.Job

func (b ByDateDesc) Len() int {
	return len(b)
}
func (b ByDateDesc) Less(i, j int) bool {
	iTime, _ := time.Parse(time.RFC822, b[i].DateAdded)
	jTime, _ := time.Parse(time.RFC822, b[j].DateAdded)

	return iTime.After(jTime)
}
func (b ByDateDesc) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func SearchJobs(db *sql.DB) {
	u, err := database.FindUser(db)

	if err != nil {
		panic(err)
	}

	var jobsToSend []sources.Job
	for _, s := range sources.All() {
		for _, tag := range u.Tags() {
			searchedJobs := sources.SearchFor(tag, s)

			if s.Name() == "berlinstartupjobs" {
				sort.Sort(ByDateDesc(searchedJobs))
			}

			if len(searchedJobs) == 0 {
				continue
			}

			//check against current lock job_history
			if history, ok := database.FindBySourceAndTag(db, s.Name(), tag); ok {
				for _, j := range searchedJobs {
					historyDate := history.MostRecent.DateAdded
					historyTitle := history.MostRecent.Title

					historyTime, err := time.Parse(time.RFC822, historyDate)
					if err != nil {
						panic(err)
					}

					jobTime, err := time.Parse(time.RFC822, j.DateAdded)

					if err != nil {
						panic(err)
					}

					if !jobTime.After(historyTime) {
						continue
					}

					if jobTime.Equal(historyTime) && j.Title == historyTitle {
						continue
					}

					jobsToSend = append(jobsToSend, j)
				}
			} else {
				jobsToSend = append(jobsToSend, searchedJobs...)
			}

			//lock history
			h := database.JobHistory{SourceName: s.Name(), Tag: tag}
			h.MostRecent = searchedJobs[0]
			database.InsertOrUpdate(db, h)
		}
	}

	if len(jobsToSend) > 0 {
		sendJobs(jobsToSend)
	}

	//todo improve algorithm for searching(especially golang:keyword in title, keyword in tags, keyword in text)
	//todo write documentation
	//todo final code review
	//todo maybe add more than 1 row per source-tag in job_history
	//todo search for remaining todos
	//todo check for _ errors

	//foreach term
	//	drop jobs with dateAdded < last date added
	//	if date == last date added drop the one with the same name
	// 	send all remaining
}

func sendJobs(jobs []sources.Job) {
	var body string

	body = "Hey, <strong>jobber</strong> found new jobs for you <br />"

	for _, j := range jobs {
		body += fmt.Sprintf("<a href=\"%s\">%s</a><br />", j.Url, j.Title)
	}
	email.Send("andrey.boar@gmail.com", "New jobs from jobber", body)
}
