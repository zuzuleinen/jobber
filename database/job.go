package database

import (
	"database/sql"
	"github.com/zuzuleinen/jobber/sources"
)

func CreateJobsTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS jobs(Title TEXT, Tag TEXT, Url TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func Save(db *sql.DB, jobs []sources.Job) {
	sql_additem := `INSERT INTO jobs(Title,Tag,Url) values(?,?,?);`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range jobs {
		_, err2 := stmt.Exec(item.Title, item.Tag, item.Url)
		if err2 != nil {
			panic(err2)
		}
	}
}

func Read(db *sql.DB) []sources.Job {
	sql_readAll := `SELECT Title, Tag, Url FROM jobs;`

	rows, err := db.Query(sql_readAll)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []sources.Job
	for rows.Next() {
		item := sources.Job{}
		err2 := rows.Scan(&item.Title, &item.Tag, &item.Url)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}
