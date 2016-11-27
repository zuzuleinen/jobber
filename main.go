package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/jobber/commands"
	"github.com/zuzuleinen/jobber/database"
	"os"
)

func main() {
	db := database.Connect()
	defer db.Close()

	if len(os.Args) < 2 {
		fmt.Println("Use `init` or `search` commands.")
		return
	}

	command := os.Args[1]
	switch command {
	case "init":
		database.CreateJobsTable(db)
		database.CreateUserTable(db)
		database.CreateJobHistoryTable(db)
		commands.SaveData(db)
		break
	case "search":
		commands.SearchJobs(db)
	}
}
