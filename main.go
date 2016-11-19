package main

import (
	"fmt"
	"github.com/zuzuleinen/jobber/sources"
)

func main() {
	j := make([]sources.Job, 0)
	for _, s := range sources.All() {
		j = append(j, s.SearchFor("php")...)
	}

	for k, v := range j {
		fmt.Println(k, v.Title, v.Url)
	}
}
