package models_group

import (
	"fmt"
	"social-network/pkg/db/models"
)

func (db *GroupDB) InsertGroup(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
		admin_id : int,
		title : string,
		about : string,
		}
	*/
	stmt := "INSERT INTO groups (admin_id, title, about) VALUES (?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["admin_id"], obj["title"], obj["about"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	newGroupId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectGroupById(map[string]any{"id": newGroupId})
}

func (db *GroupDB) SelectGroupById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, admin_id, title, about, date_creation FROM groups WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	group := models.Group{}
	err := result.Scan(&group.Id, &group.AdminId, &group.Title, &group.About, &group.DateCreation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: group}, nil
}
