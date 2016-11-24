package main

import (
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/zuzuleinen/jobber/sources"
	"github.com/zuzuleinen/jobber/database"
)

func main() {
	db := database.Connect()
	defer db.Close()

	jobs := make([]sources.Job, 0)
	for _, s := range sources.All() {
		jobs = append(jobs, sources.SearchFor("php", s)...)
	}

	for k, v := range jobs {
		fmt.Println(k, v.Title, v.Url)
	}
}
