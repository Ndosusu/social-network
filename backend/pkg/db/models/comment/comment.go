package models_comment

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *CommentDB) InsertComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			author_id : int,
			post_id : int,
			message : string,
			image : string,
		}
	*/
	stmt := "INSERT INTO comments (author_id, post_id, message, image, group_id, date) VALUES (?, ?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["post_id"], obj["message"], obj["image"], obj["group_id"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newCommentId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectCommentById(map[string]any{"id": newCommentId})
}

func (db *CommentDB) SelectCommentById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			comment_id : int,
		}
	*/
	stmt := "SELECT id, author_id, post_id, message, image, date, group_id FROM comments WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["comment_id"])

	comment := models.Comment{
		Author: &models.User{Id: utils.NOT_SCANNED},
		Post:   &models.Post{Id: utils.NOT_SCANNED},
		Image:  nil,
	}
	err := result.Scan(&comment.Id, &comment.Author.Id, &comment.Post.Id, &comment.Message, &comment.Image, &comment.Date)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: comment}, nil
}

func (db *CommentDB) DeleteComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			comment_id : int,
		}
	*/
	stmt := "DELETE FROM comments WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["comment_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
