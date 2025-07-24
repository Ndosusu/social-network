package models_rel

import (
	"fmt"
	"social-network/pkg/db/models"
)

// Queries about event membership relations
func (db *RelDB) InsertEventMemberRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			event_id : int,
			member_id : int,
		}
	*/
	stmt := "INSERT INTO event_member_rel (event_id, member_id) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["event_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
func (db *RelDB) DeleteEventMemberRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			event_id : int,
			member_id : int,
		}
	*/
	stmt := "DELETE FROM event_member_rel WHERE event_id = ? AND member_id = ?;"
	_, err := db.Conn.Exec(stmt, obj["event_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
