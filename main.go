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

	show := false
	if len(os.Args) > 2 {
		if os.Args[2] == "--show" || os.Args[2] == "-s" {
			show = true
		}
	}

	switch command {
	case "install":
		database.Remove()
		database.CreateUserTable(db)
		database.CreateMailgunTable(db)
		database.CreateJobHistoryTable(db)
		commands.SaveData(db)
		break
	case "uninstall":
		database.Remove()
		break
	case "search":
		commands.SearchJobs(db, show)
		break
	case "--help", "-h":
		commands.Help()
	default:
		commands.Help()
	}
}
