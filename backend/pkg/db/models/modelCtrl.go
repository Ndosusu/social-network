package models

import (
	"database/sql"
	"fmt"
	"social-network/config"

	_ "github.com/mattn/go-sqlite3"
)

// Struct to contain any type that result from database operations
type Response struct {
	Result any
}

// Struct to manage database connection
type DB struct {
	Conn *sql.DB
}

func (db *DB) OpenConn() {
	conn, err := sql.Open("sqlite3", config.DBPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Conn = conn
}

func (db *DB) CloseConn() {
	err := db.Conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
