package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectDB(dataSourceName string) *sql.DB {
    var err error
    db, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    return db
}


func GetDB() *sql.DB {
    return db
}


func CloseDB() {
    if db != nil {
        db.Close()
    }
}