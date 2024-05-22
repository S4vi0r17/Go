package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:benites1234@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Error pinging database: %q", err)
		}

		log.Println("Successfully connected!")
	})
}

// Return a unique instance of the database
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	var ns sql.NullString
	if s == "" {
		ns.Valid = false
	} else {
		ns.String = s
		ns.Valid = true
	}
	return ns
}
