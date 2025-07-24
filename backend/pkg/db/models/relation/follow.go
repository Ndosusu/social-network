package models_rel

import (
	"fmt"
	"social-network/pkg/db/models"
)

// Queries about follow relations
func (db *RelDB) InsertFollowRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			user_from : int,
			user_to : int,
		}
	*/
	stmt := "INSERT INTO follow_rel (user_to, user_from) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["user_to"], obj["user_from"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
func (db *RelDB) DeleteFollowRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			user_from : int,
			user_to : int,
		}
	*/
	stmt := "DELETE FROM follow_rel WHERE user_to = ? AND user_from = ?;"
	_, err := db.Conn.Exec(stmt, obj["user_to"], obj["user_from"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
