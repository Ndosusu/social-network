package models_event

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *EventDB) InsertEvent(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			group_id : int,
			title : string,
			about : string,
			date_schedule : string,
		}
	*/
	stmt := "INSERT INTO events (group_id, title, about, date_schedule, date_creation) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["group_id"], obj["title"], obj["about"], obj["date_schedule"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newEventId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectEventById(map[string]any{"event_id": newEventId})
}

func (db *EventDB) SelectEventById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			event_id : int,
		}
	*/
	stmt := "SELECT id, group_id, title, about, date_schedule, date_creation FROM events WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["event_id"])

	event := models.Event{
		Group: &models.Group{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&event.Id, &event.Group.Id, &event.Title, &event.About, &event.DateSchedule, &event.DateCreation)
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: event}, nil
}
func (db *EventDB) DeleteEvent(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			event_id : int,
		}
	*/
	stmt := "DELETE FROM events WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["event_id"])
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
