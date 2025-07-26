package models_post

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *PostDB) InsertLike(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}
	newLikeId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectLikeById(map[string]any{"id": newLikeId})
}

func (db *PostDB) SelectLikeById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, user_id, post_id, comment_id FROM likes WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	like := models.Like{
		Author:  &models.User{Id: utils.NOT_SCANNED},
		Post:    &models.Post{Id: utils.NOT_SCANNED},
		Comment: &models.Comment{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&like.Id, &like.Author.Id, &like.Post.Id, &like.Comment.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: like}, nil
}

func (db *PostDB) DeleteLike(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM likes WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
