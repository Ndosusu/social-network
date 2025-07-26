package models_group

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
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
	stmt := "INSERT INTO groups (admin_id, title, about, date) VALUES (?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["admin_id"], obj["title"], obj["about"], utils.GetCurrentTime())
	if err != nil {
		return nil, err
	}
	newGroupId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return db.SelectGroupById(map[string]any{"id": newGroupId})
}

func (db *GroupDB) SelectGroupById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			group_id : int,
		}
	*/
	stmt := "SELECT id, admin_id, title, about, date_creation FROM groups WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["group_id"])

	group := models.Group{
		Admin: &models.User{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&group.Id, &group.Admin.Id, &group.Title, &group.About, &group.DateCreation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: group}, nil
}

func (db *GroupDB) DeleteGroup(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			group_id : int,
		}
	*/
	stmt := "DELETE FROM groups WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["group_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &models.Response{Result: "Ok"}, nil
}
