package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateNewDB(path string) error {
	if exists(path) {
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("unable to create new db at %s; file exists and connot be deleted: %w\n", path, err)
		}
	}

	dbfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dbfile.Close()
	return nil

}

func OpenDBFile(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("error initializing new sqlite3 db: %w\n", err)
	}
	// defer db.Close()
	return db, nil

}
func exists(filepath string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	f.Close()
	return true

}
