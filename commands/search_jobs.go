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
		sendJobs(jobsToSend, db)
	} else {
		fmt.Println("no new jobs")
	}

	//todo search for remaining todos
	//todo check for _ errors
	//todo final code review
}

func sendJobs(jobs []sources.Job, db *sql.DB) {
	var body string

	body = "Hey, <strong>jobber</strong> found new jobs for you <br /><br />"

	for _, j := range jobs {
		body += fmt.Sprintf("<a href=\"%s\">%s</a><br />", j.Url, j.Title)
	}

	u, _ := database.FindUser(db)

	email.Send(u.Email, "New jobs from jobber", body, db)
}
