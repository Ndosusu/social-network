package models

import (
	"fmt"
	"social-network/pkg/utils"
)

type Event struct {
	Id           int
	GroupId      int
	Title        string
	About        string
	DateSchedule string
	DateCreation string
}

func (db *DB) InsertEvent(obj map[string]any) Response {
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
		return Response{0}
	}

	newEventId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectEventById(map[string]any{"id": newEventId})
}

func (db *DB) SelectEventById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, group_id, title, about, date_schedule, date_creation FROM events WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	event := Event{}
	err := result.Scan(&event.Id, &event.GroupId, &event.Title, &event.About, &event.DateSchedule, &event.DateCreation)
	if err != nil {
		fmt.Println(err)
		return Response{Event{}}
	}

	return Response{event}
}
