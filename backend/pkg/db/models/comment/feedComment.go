package models_comment

import (
	"fmt"
	"social-network/pkg/db/models"
)

func (db *CommentDB) SelectCommentsByPostId(obj map[string]any) (*models.Response, error) {
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
				COALESCE(ul.id, 0),
				CASE
					WHEN c.author_id = s.user_id THEN 1
					ELSE 0
				END AS is_client
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
		var nickname, message, date, avatar, image string
		var isClient bool

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
			&isClient,
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
					IsClient: isClient,
				},
			},
			LikeCount: likeCount,
		}

		if likeId > 0 {
			cf.Like = &models.Like{
				Id: likeId,
			}
		}

		if image != "" {
			cf.Comment.Image = &image
		}
		if avatar != "" {
			cf.Comment.Author.Avatar = &avatar
		}

		result = append(result, cf)
	}

	return &models.Response{Result: result}, nil
}
