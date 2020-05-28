package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	dbpath    string
	viddevice string
)

func init() {
	wd, _ := os.Getwd()

	// TODO: find the filepath the default DB and return it.
	DefaultDBPath = filepath.Join(wd, "store.db")
}

var DefaultDBPath string

type Timeclock struct {
	Monitor *facedetector
	db      *sql.DB
}

func flags() {
	flag.StringVar(&dbpath, "db", "$TIMECLOCK_DB", "path to sqlite database")
	flag.StringVar(&viddevice, "device", "", "video input device")
	flag.Parse()
}

func main() {
	log.SetPrefix("timeclock: ")
	flags()
	// starts with '$' == is an environmet variable.
	if strings.Contains(dbpath[:1], "$") {
		p := os.Getenv(dbpath)
		if p == "" {
			p = filepath.Join("", DefaultDBPath)
		}
		db, err := OpenDBFile(p)
		if err != nil {
			// we cant continue without a database connection.
			panic(err)
		}

	}
}
