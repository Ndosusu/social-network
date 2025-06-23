package models

import (
	"database/sql"
	"fmt"
	"social-network/config"

	_ "github.com/mattn/go-sqlite3"
)

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
