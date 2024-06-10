package storage

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:benites1234@tcp(localhost:3306)/godb?parseTime=true")
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

func timeToNull(t time.Time) sql.NullTime {
	var nt sql.NullTime
	if t.IsZero() {
		nt.Valid = false
	} else {
		nt.Time = t
		nt.Valid = true
	}
	return nt
}
