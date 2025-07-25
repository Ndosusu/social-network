package models_rel

import (
	"fmt"
	"social-network/pkg/db/models"
)

// Queries about post privacy relations
func (db *RelDB) InsertPrivacyPostRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			post_id : int,
			follower_id : int,
		}
	*/
	stmt := "INSERT INTO privacy_post_rel (post_id, follower_id) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["post_id"], obj["follower_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
func (db *RelDB) DeletePrivacyPostRel(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			post_id : int,
			follower_id : int,
		}
	*/
	stmt := "DELETE FROM privacy_post_rel WHERE post_id = ? AND follower_id = ?;"
	_, err := db.Conn.Exec(stmt, obj["post_id"], obj["follower_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
