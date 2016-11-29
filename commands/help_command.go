package commands

import "fmt"

func Help() {
	doc := `
Usage:
    jobber <command> [option]

List of commands:
  install:        Interactive project install
  uninstall:      Remove sqlite database from homer directory
  search:         Search for new jobs and send e-mail if any
  credential:     Add a new credential record.
  credentials:    List all credentials.


Options:
  -h --help         Show this screen.
`

	fmt.Println(doc)
}

