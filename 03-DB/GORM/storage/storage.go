package storage

import (
	// "database/sql"

	"log"
	"sync"

	// "time"

	// Old import
	// "github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"

	// New import
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// func NewPostgresDB() {
// 	once.Do(func() {
// 		var err error
// 		db, err = gorm.Open("postgres", "postgres://postgres:benites1234@localhost:5432/godb?sslmode=disable")
// 		if err != nil {
// 			log.Fatalf("Error opening database: %q", err)
// 		}

// 		log.Println("Successfully connected!")
// 	})
// }

func NewPostgresDB() {
	once.Do(func() {
		var err error
		dsn := "postgres://postgres:benites1234@localhost:5432/godb?sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		log.Println("Successfully connected!")
	})
}

// Return a unique instance of the database
func Pool() *gorm.DB {
	return db
}

// func stringToNull(s string) sql.NullString {
// 	var ns sql.NullString
// 	if s == "" {
// 		ns.Valid = false
// 	} else {
// 		ns.String = s
// 		ns.Valid = true
// 	}
// 	return ns
// }

// func timeToNull(t time.Time) sql.NullTime {
// 	var nt sql.NullTime
// 	if t.IsZero() {
// 		nt.Valid = false
// 	} else {
// 		nt.Time = t
// 		nt.Valid = true
// 	}
// 	return nt
// }
