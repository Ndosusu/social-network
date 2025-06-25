package models

import (
	"fmt"
	"social-network/pkg/utils"
)

type Log struct {
	Id       int
	ChatId   int
	AuthorId int
	Message  string
	Date     string
}

func (db *DB) InsertLog(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			chat_id : int,
			author_id : int,
			log : string,
		}
	*/
	stmt := "INSERT INTO chat_log (chat_id, author_id, log, date) VALUES (?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["chat_id"], obj["author_id"], obj["log"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newLogId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectLogById(map[string]any{"id": newLogId})
}

func (db *DB) SelectLogById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, chat_id, author_id, log, date FROM chat_log WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	log := Log{}
	err := result.Scan(&log.Id, &log.ChatId, &log.AuthorId, &log.Message, &log.Date)
	if err != nil {
		fmt.Println(err)
		return Response{Log{}}
	}

	return Response{log}
}
