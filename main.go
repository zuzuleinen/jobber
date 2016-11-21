package main

import (
	"fmt"
	"github.com/zuzuleinen/jobber/sources"
)

func main() {
	topics := make([]string, 2)
	topics = append(topics, "php", "golang")

	jobs := make([]sources.Job, 0)
	for _, topic := range topics {
		for _, s := range sources.All() {
			jobs = append(jobs, sources.SearchFor(topic, s)...)
		}
	}
	for k, v := range jobs {
		fmt.Println(k, v.Title, v.Url)
	}
}
