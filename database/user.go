package database

import (
	"database/sql"
	"errors"
	"strings"
)

type User struct {
	Email     string
	Interests string
}

var ErrUserNotFound = errors.New("No user was found in the database.")

//Get the user of the application
func FindUser(db *sql.DB) (*User, error) {
	users := Users(db)
	if (len(users) < 1) {
		return new(User), ErrUserNotFound
	}
	return &users[0], nil
}

//Get all user interests
func (u *User) Tags() []string {
	tags := strings.Split(u.Interests, ",")
	res := tags[:0]
	for _, v := range tags {
		if len(v) > 0 {
			res = append(res, v)
		}
	}
	return res
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

func Users(db *sql.DB) []User {
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
