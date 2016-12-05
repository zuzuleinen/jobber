package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/zuzuleinen/jobber/database"
	"os"
	"strings"
)

func SaveData(db *sql.DB) {
	saveUserData(db)
	saveMailgunData(db)

	fmt.Println("All done!")
}

func saveUserData(db *sql.DB) {
	var email string

	fmt.Println("Welcome to jobber, your simple toolkit to find jobs!")
	fmt.Println("But first, let's find out more about you :)")
	fmt.Print("What is your e-mail:")
	fmt.Scanln(&email)

	interests := make([]string, 0)
	r := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Give me an interest:")
		t, _, _ := r.ReadLine()
		interests = append(interests, string(t))
		if len(t) == 0 {
			break
		}
	}

	u := new(database.User)
	u.Email = email
	u.Interests = strings.Join(interests, ",")

	database.SaveUser(db, u)
}

func saveMailgunData(db *sql.DB) {
	m := new(database.MailgunData)

	fmt.Println("Almost done. I need your mailgun data")
	fmt.Print("Domain:")
	fmt.Scanln(&m.Domain)
	fmt.Print("Public key:")
	fmt.Scanln(&m.PubKey)
	fmt.Print("Private key:")
	fmt.Scanln(&m.PrivKey)

	database.SaveMailgun(db, m)
}
