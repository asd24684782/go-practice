package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(url string){
    var err error
    DB, err = sql.Open("postgres", "user=postgres password=5598 dbname=postgres sslmode=disable")
    
    if err != nil {
        panic(err)
    }
}

func Close(){
    DB.Close()
}