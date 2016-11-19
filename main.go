package main

import (
	"fmt"
	"github.com/zuzuleinen/jobber/sources"
)

func main() {
	jobs := make([]sources.Job, 0)
	for _, s := range sources.All() {
		jobs = append(jobs, sources.SearchFor("php", s)...)
	}

	for k, v := range jobs {
		fmt.Println(k, v.Title, v.Url)
	}
}
