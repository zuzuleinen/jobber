package main

import (
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/zuzuleinen/jobber/commands"
	"github.com/zuzuleinen/jobber/sources"
	"github.com/zuzuleinen/jobber/database"
	"os"
)



func main() {
	db := database.Connect()
	defer db.Close()

	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "init" {
			commands.SaveData()
		}
	}
	searchForJobs()
}

func searchForJobs() {
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
}
