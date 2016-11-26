package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/jobber/commands"
	"github.com/zuzuleinen/jobber/database"
	"os"
)

func main() {
	db := database.Connect()
	defer db.Close()

	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			database.CreateJobsTable(db)
			database.CreateUserTable(db)
			commands.SaveData(db)
		}
		if os.Args[1] == "search" {
			commands.SearchJobs()
		}
	}
}
