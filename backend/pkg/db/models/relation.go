package models

import "fmt"

func (db *DB) InsertFollowRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			user_from : int
			user_to : int,
		}
	*/
	stmt := "INSERT INTO follow_rel (user_to, user_from) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["user_to"], obj["user_from"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
