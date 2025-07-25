package models_post

import (
	"fmt"
	"social-network/pkg/db/models"
)

func (db *PostDB) GetGlobalFeed(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
			{
				session_uuid : string,
				last_id : int,
				limit : int
			}
	*/
	fmt.Println(obj)

	stmt := `SELECT 
				p.id,
				p.author_id,
				u.nick_name,
				u.avatar,
				p.message,
				COALESCE(p.image, ''),
				p.date,
				p.privacy_mode,
				COALESCE(p.group_id, 0),
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id),
				COALESCE(g.title, '')
			FROM posts p
			JOIN users u ON p.author_id = u.id
			JOIN sessions s ON s.uuid = ?
			LEFT JOIN groups g ON g.id = p.group_id
			LEFT JOIN follow_rel f ON f.user_from = s.user_id AND f.user_to = p.author_id
			LEFT JOIN group_member_rel gmr ON gmr.group_id = p.group_id AND gmr.member_id = s.user_id
			LEFT JOIN privacy_post_rel pr ON pr.post_id = p.id AND pr.follower_id = s.user_id
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE
				(
					p.author_id = s.user_id
					OR (p.privacy_mode = 0 AND u.private_mode = 0)
					OR (f.user_to IS NOT NULL AND p.privacy_mode IN (1,2))
					OR (p.privacy_mode = 3 AND pr.follower_id IS NOT NULL)
					OR (p.group_id IS NOT NULL AND gmr.group_id IS NOT NULL)
				)
				AND p.id < ?
			GROUP BY p.id
			ORDER BY p.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"], obj["last_id"], obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PostFeed
	for rows.Next() {
		var p models.Post
		var postImage string
		var author models.User
		var likeCount, commentCount int
		var groupTitle string
		var groupID int

		err := rows.Scan(
			&p.Id,
			&p.AuthorId,
			&author.Nickname,
			&author.Avatar,
			&p.Message,
			&postImage,
			&p.Date,
			&p.PrivacyMode,
			&groupID,
			&likeCount,
			&commentCount,
			&groupTitle,
		)
		if err != nil {
			return nil, err
		}

		if postImage != "" {
			p.Image = &postImage
		}
		p.GroupId = groupID
		p.Author = &author

		result = append(result, models.PostFeed{
			Post:         &p,
			LikeCount:    likeCount,
			CommentCount: commentCount,
			GroupTitle:   groupTitle,
		})
	}
	return &models.Response{Result: result}, nil
}

func (db *PostDB) GetFollowFeed(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
			{
				session_uuid : string,
				last_id : int,
				limit : int
			}
	*/

	stmt := `SELECT 
				p.id,
				p.author_id,
				u.nick_name,
				u.avatar,
				p.message,
				COALESCE(p.image, ''),
				p.date,
				p.privacy_mode,
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id),
			FROM posts p
			JOIN users u ON p.author_id = u.id
			JOIN sessions s ON s.uuid = ?
			LEFT JOIN follow_rel f ON f.user_from = s.user_id AND f.user_to = p.author_id
			LEFT JOIN privacy_post_rel pr ON pr.post_id = p.id AND pr.follower_id = s.user_id
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE
				(
					OR (f.user_to IS NOT NULL AND p.privacy_mode IN (1,2))
					OR (p.privacy_mode = 3 AND pr.follower_id IS NOT NULL)
				)
				AND p.id < ?
			GROUP BY p.id
			ORDER BY p.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"], obj["last_id"], obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PostFeed
	for rows.Next() {
		var p models.Post
		var postImage string
		var author models.User
		var likeCount, commentCount int

		err := rows.Scan(
			&p.Id,
			&p.AuthorId,
			&author.Nickname,
			&author.Avatar,
			&p.Message,
			&postImage,
			&p.Date,
			&p.PrivacyMode,
			&likeCount,
			&commentCount,
		)
		if err != nil {
			return nil, err
		}

		if postImage != "" {
			p.Image = &postImage
		}
		p.Author = &author

		result = append(result, models.PostFeed{
			Post:         &p,
			LikeCount:    likeCount,
			CommentCount: commentCount,
		})
	}
	return &models.Response{Result: result}, nil
}

func (db *PostDB) GetGroupFeed(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
			{
				group_id : int,
				last_id : int,
				limit : int
			}
	*/

	stmt := `SELECT 
				p.id,
				p.author_id,
				u.nick_name,
				u.avatar,
				p.message,
				COALESCE(p.image, ''),
				p.date,
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id),
			FROM posts p
			JOIN users u ON p.author_id = u.id
			LEFT JOIN groups g ON g.id = ? 
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE
				(
					p.group_id IS NOT NULL AND gmr.group_id IS NOT NULL
				)
				AND p.id < ?
			GROUP BY p.id
			ORDER BY p.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["group_id"], obj["last_id"], obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PostFeed
	for rows.Next() {
		var p models.Post
		var postImage string
		var author models.User
		var likeCount, commentCount int

		err := rows.Scan(
			&p.Id,
			&p.AuthorId,
			&author.Nickname,
			&author.Avatar,
			&p.Message,
			&postImage,
			&p.Date,
			&likeCount,
			&commentCount,
		)
		if err != nil {
			return nil, err
		}

		if postImage != "" {
			p.Image = &postImage
		}
		p.Author = &author

		result = append(result, models.PostFeed{
			Post:         &p,
			LikeCount:    likeCount,
			CommentCount: commentCount,
		})
	}
	return &models.Response{Result: result}, nil
}
