package commands

import (
	"fmt"
	"github.com/zuzuleinen/jobber/sources"
	"github.com/zuzuleinen/jobber/email"
	"database/sql"
	"github.com/zuzuleinen/jobber/database"
)

func SearchJobs(db *sql.DB) {
	u, err := database.FindUser(db)

	if err != nil {
		panic(err)
	}

	jobs := make([]sources.Job, 0)
	for _, topic := range u.Tags() {
		for _, s := range sources.All() {
			jobs = append(jobs, sources.SearchFor(topic, s)...)
		}
	}
	for _, j := range jobs {
		fmt.Println(j.Tag, ":", j.Title, j.Url, ":", j.DateAdded)
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
