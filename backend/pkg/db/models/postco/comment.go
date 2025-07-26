package models_post

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *PostDB) InsertComment(obj map[string]any) (*models.Response, error) {
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

func (db *PostDB) SelectCommentById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, author_id, post_id, message, image, date, group_id FROM comments WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

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

func (db *PostDB) DeleteComment(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}

func (db *PostDB) SelectCommentsByPostId(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			session_uuid : string,
			post_id : int,
			last_id : int,
			limit : int,
		}
	*/

	stmt := `SELECT
				c.id,
				c.author_id,
				u.nick_name,
				COALESCE(u.avatar,''), 
				c.message, 
				COALESCE(c.image, ''), 
				c.date, 
				COUNT(DISTINCT l.id),
				COALESCE(ul.id, 0)
			FROM comments c
			JOIN users u ON c.author_id = u.id
			JOIN sessions s ON s.uuid = ?
			LEFT JOIN likes l ON l.comment_id = c.id
			LEFT JOIN likes ul ON ul.user_id = s.user_id AND ul.comment_id = c.id
			WHERE 
				c.post_id = ? 
				AND c.id < ?
			GROUP BY c.id, c.author_id, u.nick_name, u.avatar, c.message, c.image, c.date, ul.id
			ORDER BY c.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"], obj["post_id"], obj["last_id"], obj["limit"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []models.CommentFeed
	for rows.Next() {
		var commentId, authorId, likeCount, likeId int
		var nickname, message, date string
		var avatar, image *string
		err := rows.Scan(
			&commentId,
			&authorId,
			&nickname,
			&avatar,
			&message,
			&image,
			&date,
			&likeCount,
			&likeId,
		)
		if err != nil {
			return nil, err
		}

		cf := models.CommentFeed{
			Comment: &models.Comment{
				Id:      commentId,
				Message: message,
				Date:    date,
				Author: &models.User{
					Id:       authorId,
					Nickname: nickname,
				},
			},
			Like: &models.Like{
				Id: likeId,
			},
			LikeCount: likeCount,
		}

		if image != nil && *image != "" {
			cf.Comment.Image = image
		}
		if avatar != nil && *avatar != "" {
			cf.Comment.Author.Avatar = avatar
		}

		result = append(result, cf)
	}

	return &models.Response{Result: result}, nil
}
