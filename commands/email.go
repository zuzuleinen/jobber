package commands

import (
	"fmt"
)

func SaveData() {
	var email, prompt string

	fmt.Println("Welcome to jobber, your simple toolkit to find jobs!")
	fmt.Println("But first, let's find out more about you :)")
	fmt.Print("What is your e-mail:")
	fmt.Scanln(&email)
	fmt.Println("Cool, your e-mail is", email)

	interests := make([]string, 1)
	for prompt != "." {
		fmt.Print("Give me an interest:")
		fmt.Scan(&prompt)
		if prompt != "." {
			interests = append(interests, prompt)
		}
	}
	fmt.Println(interests)
}
