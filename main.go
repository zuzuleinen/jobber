package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/jobber/commands"
	"github.com/zuzuleinen/jobber/database"
	"os"
	"strings"
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

func ParseTime(d string) string {
	if strings.Contains(d, "hours") {
		d = strings.Replace(d, "hours", "h", -1)
	}
	if strings.Contains(d, "hour") {
		d = strings.Replace(d, "hour", "h", -1)
	}
	d = strings.Replace(d, "ago", "", -1)
	d = strings.Replace(d, " ", "", -1)
	d = strings.Replace(d, "<", "", -1)
	d = "-" + d
	return d
}
