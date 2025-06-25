package models

import (
	"fmt"
	"social-network/pkg/utils"
)

type Notif struct {
	Id         int
	NotifType  int
	ReceiverId int
	SenderId   int
	GroupId    int
	EventId    int
}

func (db *DB) InsertNotif(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			type : int,
			user_to : int
			user_from : int,
			group_id : int,
			event_id : int,
		}
	*/
	stmt := "INSERT INTO notifications (type, user_to, user_from, group_id, event_id) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["group_id"], obj["title"], obj["about"], obj["date_schedule"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newNotifId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectNotifById(map[string]any{"id": newNotifId})
}

func (db *DB) SelectNotifById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, type, user_to, user_from, group_id, event_id FROM notifications WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	notif := Notif{}
	err := result.Scan(&notif.Id, &notif.NotifType, &notif.ReceiverId, &notif.SenderId, &notif.GroupId, &notif.EventId)
	if err != nil {
		fmt.Println(err)
		return Response{Notif{}}
	}

	return Response{notif}
}
func (db *DB) DeleteNotif(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM notifications WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
