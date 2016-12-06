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

func InsertOrUpdate(db *sql.DB, j JobHistory) {
	sql_select := `SELECT * FROM job_history WHERE SourceName = ? AND Tag = ?`

	rows, err := db.Query(sql_select, j.SourceName, j.Tag)
	if err != nil {
		panic(err)
	}

	if !rows.Next() {
		InsertHistories(db, []JobHistory{j})
	} else {
		rows.Close()
		Update(db, j)
	}
}

func InsertHistories(db *sql.DB, histories []JobHistory) {
	sql_insert := `INSERT INTO job_history(SourceName,Tag,LastDateAdded, LastTitle) values(?,?,?,?);`

	stmt, err := db.Prepare(sql_insert)
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

//Update job_history.LastTitle and job_history.LastDateAdded by SourceName and Tag
func Update(db *sql.DB, j JobHistory) {
	sql_update := `UPDATE job_history
	SET LastTitle = ?, LastDateAdded = ?
	WHERE SourceName = ? AND Tag = ?;`

	stmt, err := db.Prepare(sql_update)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(j.MostRecent.Title, j.MostRecent.DateAdded, j.SourceName, j.Tag)
	if err2 != nil {
		panic(err2)
	}
}

func FindBySourceAndTag(db *sql.DB, sourceName string, tag string) (*JobHistory, bool) {
	sql_select := `SELECT * FROM job_history WHERE SourceName = ? AND Tag = ?`

	rows, err := db.Query(sql_select, sourceName, tag)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	j := new(JobHistory)
	if rows.Next() {
		err2 := rows.Scan(&j.SourceName, &j.Tag, &j.MostRecent.DateAdded, &j.MostRecent.Title)
		if err2 != nil {
			panic(err2)
		}

		return j, true
	}
	return new(JobHistory), false
}
