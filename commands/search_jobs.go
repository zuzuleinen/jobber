package commands

import (
	"database/sql"
	"fmt"
	"github.com/zuzuleinen/jobber/database"
	"github.com/zuzuleinen/jobber/email"
	"github.com/zuzuleinen/jobber/sources"
)

func SearchJobs(db *sql.DB) {
	u, err := database.FindUser(db)
	debug := false

	if err != nil {
		panic(err)
	}

	jobs := make([]sources.Job, 0)
	histories := make([]database.JobHistory, 0)
	for _, s := range sources.All() {
		for _, tag := range u.Tags() {
			h := database.JobHistory{SourceName: s.Name(), Tag: tag}

			searchedJobs := sources.SearchFor(tag, s)
			jobs = append(jobs, searchedJobs...)

			h.MostRecent = searchedJobs[0]
			histories = append(histories, h)
		}
	}

	//todo standard format for db last date
	//todo deal with stackoverflow date
	//todo send jobs to e-mail
	//todo improve algorithm for searching(especially golang:keyword in title, keyword in tags, heiword in text)
	//todo check why jobs from stack overflow are not ordered

	if debug {
		for _, j := range jobs {
			fmt.Println(j.Tag, ":", j.Title, j.Url, ":", j.DateAdded)
		}
		fmt.Println("-----------------------")
		fmt.Println(histories)
	}

	for _, h := range histories {
		database.InsertOrUpdate(db, h)
	}

	//foreach term
	//	drop jobs with dateAdded < last date added
	//	if date == last date added drop the one with the same name
	// 	send all remaining
}

func sendJobs(jobs []sources.Job) {
	var body string

	body = "Hey, <strong>jobber</strong> found new jobs for you <br />"

	for _, j := range jobs {
		body += fmt.Sprintf("<a href=\"%s\">%s</a>", j.Url, j.Title)
	}
	email.Send("andrey.boar@gmail.com", "New jobs from jobber", body)
}
