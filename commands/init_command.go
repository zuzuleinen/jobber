package commands

import (
	"fmt"
	"bufio"
	"os"
)

func SaveData() {
	var email string

	fmt.Println("Welcome to jobber, your simple toolkit to find jobs!")
	fmt.Println("But first, let's find out more about you :)")
	fmt.Print("What is your e-mail:")
	fmt.Scanln(&email)
	fmt.Println("Cool, your e-mail is", email)

	interests := make([]string, 1)
	r := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Give me an interest:")
		t, _, _ := r.ReadLine()
		interests = append(interests, string(t))
		if len(t) == 0 {
			break;
		}
	}
	fmt.Println(interests)
}
