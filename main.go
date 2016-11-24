package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/jobber/commands"
	"github.com/zuzuleinen/jobber/database"
	"github.com/zuzuleinen/jobber/sources"
	"os"
)

func main() {
	db := database.Connect()
	defer db.Close()

	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			commands.SaveData()
			database.CreateJobsTable(db)
		}
	} else {
		searchForJobs(db)
	}
}

func searchForJobs(db *sql.DB) {
	topics := make([]string, 0)
	topics = append(topics, "php", "golang")

	jobs := make([]sources.Job, 0)
	for _, topic := range topics {
		for _, s := range sources.All() {
			jobs = append(jobs, sources.SearchFor(topic, s)...)
		}
	}
	for _, j := range jobs {
		fmt.Println(j.Tag, ":", j.Title, j.Url)
	}
	database.Save(db, jobs)
}
