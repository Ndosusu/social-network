package models

import (
	"fmt"
)

type Chat struct {
	Id         int
	ReceiverId int
	SenderId   int
	GroupId    int
}

func (db *DB) InsertChat(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			user_to : int
			user_from : int,
			group_id : int,
		}
	*/
	stmt := "INSERT INTO chats (user_to, user_from, group_id) VALUES (?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["user_to"], obj["user_from"], obj["group_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newChatId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectChatById(map[string]any{"id": newChatId})
}

func (db *DB) SelectChatById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, user_to, user_from, group_id FROM chats WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	chat := Chat{}
	err := result.Scan(&chat.Id, &chat.ReceiverId, &chat.SenderId, &chat.GroupId)
	if err != nil {
		fmt.Println(err)
		return Response{Chat{}}
	}

	return Response{chat}
}
