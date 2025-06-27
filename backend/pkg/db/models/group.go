package models

import "fmt"

type Group struct {
	Id           int
	AdminId      int
	Title        string
	About        string
	DateCreation string
}

func (db *DB) InsertGroup(obj map[string]any) Response {
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
		return Response{0}
	}

	newGroupId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectGroupById(map[string]any{"id": newGroupId})
}

func (db *DB) SelectGroupById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, admin_id, title, about, date_creation FROM groups WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	group := Group{}
	err := result.Scan(&group.Id, &group.AdminId, &group.Title, &group.About, &group.DateCreation)
	if err != nil {
		fmt.Println(err)
		return Response{Group{}}
	}

	return Response{group}
}
