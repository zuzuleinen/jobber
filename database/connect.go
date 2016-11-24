package database

import (
	"database/sql"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		panic(err.Error())
	}
	if db == nil {
		panic("db nil")
	}
	return db
}
