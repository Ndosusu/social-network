package models_notif

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *NotifDB) InsertNotif(obj map[string]any) (*models.Response, error) {
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
	stmt := "INSERT INTO notifications (type, user_to, user_from, group_id, event_id, date_creation) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["type"], obj["user_to"], obj["user_from"], obj["group_id"], obj["event_id"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newNotifId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectNotifById(map[string]any{"id": newNotifId})
}

func (db *NotifDB) SelectNotifById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, type, user_to, user_from, group_id, event_id, date_creation FROM notifications WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	notif := models.Notif{}
	err := result.Scan(&notif.Id, &notif.NotifType, &notif.ReceiverId, &notif.SenderId, &notif.GroupId, &notif.EventId, &notif.DateCreation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: notif}, nil
}
func (db *NotifDB) DeleteNotif(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
