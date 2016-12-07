package database

import (
	"os"
)

func Remove() {
	err := os.Remove(DbPath())
	if err != nil {
		panic(err)
	}
}
