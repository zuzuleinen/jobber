package database

import "database/sql"

type MailgunData struct {
	PubKey string
	PrivKey string
}

func CreateMailgunTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS mailgun(PubKey TEXT, PrivKey TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
func SaveMailgun(db *sql.DB, m *MailgunData) {
	sql_additem := `INSERT INTO mailgun(PubKey,PrivKey) values(?,?);`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(m.PubKey, m.PrivKey)
	if err2 != nil {
		panic(err2)
	}
}