package models_group

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *GroupDB) InsertEvent(obj map[string]any) (*models.Response, error) {
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
	return db.SelectEventById(map[string]any{"id": newEventId})
}

func (db *GroupDB) SelectEventById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, group_id, title, about, date_schedule, date_creation FROM events WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	event := models.Event{}
	err := result.Scan(&event.Id, &event.GroupId, &event.Title, &event.About, &event.DateSchedule, &event.DateCreation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: event}, nil
}
func (db *GroupDB) DeleteEvent(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM events WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
