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

	command := "--help"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "install":
		removeDatabase()
		database.CreateUserTable(db)
		database.CreateJobHistoryTable(db)
		commands.SaveData(db)
		break
	case "uninstall":
		removeDatabase()
		break
	case "search":
		commands.SearchJobs(db)
		break
	case "--help", "-h":
		commands.Help()
	default:
		commands.Help()
	}
}

func removeDatabase() {
	err := os.Remove(database.DbPath())
	if err != nil {
		panic(err)
	}
}
