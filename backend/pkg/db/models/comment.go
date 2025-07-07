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

type CommentWithAuthor struct {
	Id           int    `json:"id"`
	AuthorId     int    `json:"author_id"`
	PostId       int    `json:"post_id"`
	Message      string `json:"message"`
	Image        string `json:"image"`
	Date         string `json:"date"`
	GroupId      int    `json:"group_id"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
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
	stmt := "INSERT INTO comments (author_id, post_id, message, image, group_id, date) VALUES (?, ?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["post_id"], obj["message"], obj["image"], obj["group_id"], utils.GetCurrentTime())
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

func (db *DB) DeleteComment(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM comments WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}

func (db *DB) SelectCommentsByPostId(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			post_id : int,
		}
	*/
	stmt := "SELECT id, author_id, post_id, message, image, date, group_id FROM comments WHERE post_id = ? ORDER BY date ASC;"
	rows, err := db.Conn.Query(stmt, obj["post_id"])
	if err != nil {
		fmt.Println(err)
		return Response{[]Comment{}}
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		comment := Comment{}
		err := rows.Scan(&comment.Id, &comment.AuthorId, &comment.PostId, &comment.Message, &comment.Image, &comment.Date, &comment.GroupId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		comments = append(comments, comment)
	}

	return Response{comments}
}

func (db *DB) SelectCommentsByPostIdWithAuthor(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			post_id : int,
		}
	*/
	stmt := `SELECT 
		c.id, c.author_id, c.post_id, c.message, c.image, c.date, c.group_id,
		u.first_name, u.last_name, u.avatar
		FROM comments c
		JOIN users u ON c.author_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.date ASC;`

	rows, err := db.Conn.Query(stmt, obj["post_id"])
	if err != nil {
		fmt.Println(err)
		return Response{[]CommentWithAuthor{}}
	}
	defer rows.Close()

	var comments []CommentWithAuthor
	for rows.Next() {
		var comment CommentWithAuthor
		var firstName, lastName string
		err := rows.Scan(&comment.Id, &comment.AuthorId, &comment.PostId, &comment.Message, &comment.Image, &comment.Date, &comment.GroupId, &firstName, &lastName, &comment.AuthorAvatar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		comment.AuthorName = firstName + " " + lastName
		comments = append(comments, comment)
	}

	return Response{comments}
}
