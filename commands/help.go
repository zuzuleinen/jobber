package commands

import "fmt"

func Help() {
	doc := `
Usage:
    jobber <command> [option]
    jobber search -s

List of commands:
  install:        Interactive project install
  uninstall:      Remove sqlite database from home directory
  search:         Search for new jobs and send e-mail if any


Options:
  -h --help         Show this screen.
  -s --show         Display output when using 'jobber search'
`

	fmt.Println(doc)
}
