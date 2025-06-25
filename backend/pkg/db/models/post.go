package models

import (
	"fmt"
	"social-network/pkg/utils"
)

type Post struct {
	Id          int
	AuthorId    int
	Message     string
	Image       string
	Date        string
	PrivacyMode int
	GroupId     int
}

func (db *DB) InsertPost(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			author_id : int,
			message : string,
			image : string,
			privacy_mode : int,
			group_id : int,
		}
	*/
	stmt := "INSERT INTO posts (author_id, message, image, privacy_mode, group_id, date) VALUES (?, ?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["message"], obj["image"], obj["privacy_mode"], obj["group_id"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newPostId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectPostById(map[string]any{"id": newPostId})
}

func (db *DB) SelectPostById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	post := Post{}
	err := result.Scan(&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId)
	if err != nil {
		fmt.Println(err)
		return Response{Post{}}
	}

	return Response{post}
}
func (db *DB) DeletePost(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM posts WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
