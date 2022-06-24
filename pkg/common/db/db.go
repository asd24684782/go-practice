package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Init(url string) *sql.DB {
    db, err := sql.Open("postgres", "user=postgres password=5598 dbname=postgres sslmode=disable")
    
    if err != nil {
        panic(err)
    }
    return db
}

func Close(db *sql.DB){
    db.Close()
}