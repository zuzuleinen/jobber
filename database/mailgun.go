package database

import "database/sql"

type MailgunData struct {
	Domain  string
	PubKey  string
	PrivKey string
}

func CreateMailgunTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS mailgun(Domain TEXT, PubKey TEXT, PrivKey TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
func SaveMailgun(db *sql.DB, m *MailgunData) {
	sql_additem := `INSERT INTO mailgun(Domain,PubKey,PrivKey) values(?,?,?);`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(m.Domain, m.PubKey, m.PrivKey)
	if err2 != nil {
		panic(err2)
	}
}

func Mailgun(db *sql.DB) *MailgunData {
	sql_select := `SELECT * FROM mailgun LIMIT 1`

	rows, err := db.Query(sql_select)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	m := new(MailgunData)
	if rows.Next() {
		err2 := rows.Scan(&m.Domain, &m.PubKey, &m.PrivKey)
		if err2 != nil {
			panic(err2)
		}

		return m
	}
	return new(MailgunData)
}