package database

import (
	"github.com/mitchellh/go-homedir"
	"os"
)

const DB_FILE = "jobber.db"

func dbPath() string {
	homeDir, _ := homedir.Dir()
	return homeDir + string(os.PathSeparator) + DB_FILE
}
