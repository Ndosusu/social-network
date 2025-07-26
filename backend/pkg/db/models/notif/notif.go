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

// GROUP OR EVENT NEED TO MODIFY LIKE I DID FOR LIKE
func (db *NotifDB) SelectNotifById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			notif_id : int,
		}
	*/
	stmt := "SELECT id, type, user_to, user_from, group_id, event_id, date_creation FROM notifications WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["notif_id"])

	notif := models.Notif{
		Receiver: &models.User{Id: utils.NOT_SCANNED},
		Sender:   &models.User{Id: utils.NOT_SCANNED},
		Group:    &models.Group{Id: utils.NOT_SCANNED},
		Event:    &models.Event{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&notif.Id, &notif.NotifType, &notif.Receiver.Id, &notif.Sender.Id, &notif.Group.Id, &notif.Event.Id, &notif.DateCreation)
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
			notif_id : int,
		}
	*/
	stmt := "DELETE FROM notifications WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["notif_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
