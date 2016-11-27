package database

import (
	"database/sql"
	"github.com/zuzuleinen/jobber/sources"
)

type JobHistory struct {
	SourceName string
	Tag        string
	MostRecent sources.Job
}

func CreateJobHistoryTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS job_history(SourceName TEXT, Tag TEXT, LastDateAdded TEXT, LastTitle TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func SaveHistories(db *sql.DB, histories []JobHistory) {
	sql_additem := `INSERT INTO job_history(SourceName,Tag,LastDateAdded, LastTitle) values(?,?,?,?);`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range histories {
		_, err2 := stmt.Exec(item.SourceName, item.Tag, item.MostRecent.DateAdded, item.MostRecent.Title)
		if err2 != nil {
			panic(err2)
		}
	}
}

func Histories(db *sql.DB) []JobHistory {
	sql_readAll := `SELECT SourceName, Tag, LastDateAdded, LastTitle FROM job_history;`

	rows, err := db.Query(sql_readAll)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []JobHistory
	for rows.Next() {
		item := JobHistory{}
		err2 := rows.Scan(&item.SourceName, &item.MostRecent.Tag, &item.MostRecent.Title, &item.MostRecent.DateAdded)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

