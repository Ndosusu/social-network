package models_chat

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *ChatDB) InsertLog(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	newLogId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectLogById(map[string]any{"id": newLogId})
}

func (db *ChatDB) SelectLogById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			log_id : int,
		}
	*/
	stmt := "SELECT id, chat_id, author_id, log, date FROM chat_log WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["log_id"])

	log := models.Log{
		Author: &models.User{Id: utils.NOT_SCANNED},
		Chat:   &models.Chat{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&log.Id, &log.Chat.Id, &log.Author.Id, &log.Message, &log.Date)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: log}, nil
}
func (db *ChatDB) DeleteLog(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			log_id : int,
		}
	*/
	stmt := "DELETE FROM chat_log WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["log_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
