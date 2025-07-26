package models_chat

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *ChatDB) InsertChat(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	newChatId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectChatById(map[string]any{"id": newChatId})
}

func (db *ChatDB) SelectChatById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, user_to, user_from, group_id FROM chats WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	chat := models.Chat{
		Receiver: &models.User{Id: utils.NOT_SCANNED},
		Sender:   &models.User{Id: utils.NOT_SCANNED},
		Group:    &models.Group{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&chat.Id, &chat.Receiver.Id, &chat.Sender.Id, &chat.Group.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: chat}, nil
}

func (db *ChatDB) DeleteChat(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{

			id : int,
		}
	*/
	stmt := "DELETE FROM chat WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
