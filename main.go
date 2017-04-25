package main

import (
	"database/sql"
	"os"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("POSTGRES_URL")
	if url == "" {
		log.Fatal("Must supply postgres URL to connect to as environment variable POSTGRES_URL")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE users (
	   id SERIAL PRIMARY KEY,
	   name TEXT);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`INSERT INTO users (name)
		 VALUES ('Swatson');`)
	if err != nil {
		log.Fatal(err)
	}
}
