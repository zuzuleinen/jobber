package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/jobber/commands"
	"github.com/zuzuleinen/jobber/database"
	"os"
	"time"
	"strings"
)

func main() {
	db := database.Connect()
	defer db.Close()

	dateAdded := "2 hour ago"
	fmt.Println(getTime(dateAdded))

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

func getTime(d string) time.Time {
	if strings.Contains(d, "hours") {
		d = strings.Replace(d, "hours", "h", -1)
	}
	if strings.Contains(d, "hour") {
		d = strings.Replace(d, "hour", "h", -1)
	}

	//max unit is h so this bellow needs to be refactor
	if strings.Contains(d, "yesterday") {
		d = strings.Replace(d, "yesterday", "1d", -1)
	}
	if strings.Contains(d, "days") {
		d = strings.Replace(d, "days", "d", -1)
	}
	if strings.Contains(d, "week") {
		d = strings.Replace(d, "days", "d", -1)
	}

	d = strings.Replace(d, "ago", "", -1)
	d = strings.Replace(d, " ", "", -1)
	d = "-" + d
	fmt.Println(d)
	duration, _ := time.ParseDuration(d)

	return time.Now().Add(duration)
}
