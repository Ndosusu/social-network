package models

import (
	"database/sql"
	"fmt"
	"social-network/config"
	"social-network/pkg/db/models/types"

	_ "github.com/mattn/go-sqlite3"
)

type (
	User     = types.User
	Session  = types.Session
	Post     = types.Post
	PostFeed = types.PostFeed
	Comment  = types.Comment
	Like     = types.Like
	Notif    = types.Notif
	Group    = types.Group
	Event    = types.Event
	Chat     = types.Chat
	Log      = types.Log
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
	conn, err := sql.Open("sqlite3", config.DBPath+config.DBName)
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
