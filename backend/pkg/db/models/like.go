package models

import (
	"fmt"
)

type Like struct {
	Id        int
	UserId    int
	PostId    int
	CommentId int
}

func (db *DB) InsertLike(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			user_id : int
			post_id : int,
			comment_id : int,
		}
	*/
	stmt := "INSERT INTO likes (user_id, post_id, comment_id) VALUES (?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["user_id"], obj["post_id"], obj["comment_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newLikeId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectLikeById(map[string]any{"id": newLikeId})
}

func (db *DB) SelectLikeById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, user_id, post_id, comment_id FROM likes WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	like := Like{}
	err := result.Scan(&like.Id, &like.UserId, &like.PostId, &like.CommentId)
	if err != nil {
		fmt.Println(err)
		return Response{Like{}}
	}

	return Response{like}
}
