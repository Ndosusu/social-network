package models_rel

import (
	"fmt"
	"social-network/pkg/db/models"
)

// Queries about group membership relations
func (db *RelDB) InsertGroupMemberRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			group_id : int,
			member_id : int,
		}
	*/
	stmt := "INSERT INTO group_member_rel (group_id, member_id) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["group_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
func (db *RelDB) DeleteGroupMemberRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			group_id : int,
			member_id : int,
		}
	*/
	stmt := "DELETE FROM group_member_rel WHERE group_id = ? AND member_id = ?;"
	_, err := db.Conn.Exec(stmt, obj["group_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
