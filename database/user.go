package database

import (
	"database/sql"
)

type User struct {
	Email     string
	Interests string
}

func CreateUserTable(db *sql.DB) {
	// create table if not exists
	sql_table := `CREATE TABLE IF NOT EXISTS users(Email TEXT, Interests TEXT);`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func SaveUser(db *sql.DB, u *User) {
	sql_additem := `INSERT INTO users(Email,Interests) values(?,?);`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(u.Email, u.Interests)
	if err2 != nil {
		panic(err2)
	}

}

func ReadUser(db *sql.DB) []User {
	sql_readAll := `SELECT Email, Interests FROM users;`

	rows, err := db.Query(sql_readAll)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		item := User{}
		err2 := rows.Scan(&item.Email, &item.Interests)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}
