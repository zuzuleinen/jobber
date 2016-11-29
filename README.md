# jobber

                                           _       _     _               
                                          (_)     | |   | |              
                                           _  ___ | |__ | |__   ___ _ __ 
                                          | |/ _ \| '_ \| '_ \ / _ \ '__|
                                          | | (_) | |_) | |_) |  __/ |   
                                          | |\___/|_.__/|_.__/ \___|_|   
                                         _/ |                            
                                        |__/                             




## About

**jobber** is a CLI tool for searching tech jobs in Berlin area. Stop browsing and let jobber do the work for you.
After install, the tool will search for jobs on [StackOverflow](http://stackoverflow.com/) and [berlinstartupjobs.com](http://berlinstartupjobs.com/) based on your preferences. Once it finds a new job it will send you an e-mail.

## Requirements

Go should be [installed and set up](https://golang.org/doc/install) on your system. Tested with version **go1.7.1**

[SQLite](https://sqlite.org/) must be installed on your system. This version was tested with **sqlite3**.

You should have a [mailgun](http://www.mailgun.com/) account. You can set-up a free account there and just use the sandbox credentials. It should be enough for the jobs you'll get in your inbox(max 10.000 mails per month).

## Installation

* Prepare the executable 

```shell
$ go get github.com/zuzuleinen/jobber
$ cd $GOPATH/src/github.com/zuzuleinen/jobber/
$ go install
$ jobber install
$ jobber
```

## Usage

```shell
Usage:
    jobber <command> [option]

List of commands:
  install:        Interactive project install
  uninstall:      Remove sqlite database from homer directory
  search:         Search for new jobs and send e-mail if any


Options:
  -h --help         Show this screen.
```



## Questions or suggestions
If you encounter a problem feel free to [open](https://github.com/zuzuleinen/dave/issues/new) an issue.
