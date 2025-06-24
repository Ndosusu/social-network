package models

import (
	"fmt"
	"social-network/pkg/utils"
)

type Comment struct {
	Id       int
	AuthorId int
	PostId   int
	Message  string
	Image    string
	Date     string
	GroupId  int
}

func (db *DB) InsertComment(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			author_id : int,
			post_id : int,
			message : string,
			image : string,
			privacy_mode : int,
			group_id : int,
		}
	*/
	stmt := "INSERT INTO comments (author_id, post_id, message, image, group_id, date) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["message"], obj["image"], obj["privacy_mode"], obj["group_id"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newCommentId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectCommentById(map[string]any{"id": newCommentId})
}

func (db *DB) SelectCommentById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, author_id, post_id, message, image, date, group_id FROM comments WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	comment := Comment{}
	err := result.Scan(&comment.Id, &comment.AuthorId, &comment.PostId, &comment.Message, &comment.Image, &comment.Date, &comment.GroupId)
	if err != nil {
		fmt.Println(err)
		return Response{Comment{}}
	}

	return Response{comment}
}
