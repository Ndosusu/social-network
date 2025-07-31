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
				COALESCE(u.nick_name, ''),
				u.first_name,
				u.last_name,
				COALESCE(u.avatar, ''),
				p.message,
				COALESCE(p.image, ''),
				p.date,
				p.privacy_mode,
				COALESCE(p.group_id, 0),
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id),
				COALESCE(g.title, ''),
				COALESCE(ul.id, 0),
				CASE
					WHEN p.author_id = s.user_id THEN 1
					ELSE 0
				END AS is_client
			FROM posts p
			JOIN users u ON p.author_id = u.id
			JOIN sessions s ON s.uuid = ?
			LEFT JOIN groups g ON g.id = p.group_id
			LEFT JOIN follow_rel f ON f.user_from = s.user_id AND f.user_to = p.author_id
			LEFT JOIN group_member_rel gmr ON gmr.group_id = p.group_id AND gmr.member_id = s.user_id
			LEFT JOIN privacy_post_rel pr ON pr.post_id = p.id AND pr.follower_id = s.user_id
			LEFT JOIN likes l ON l.post_id = p.id
        	LEFT JOIN likes ul ON ul.post_id = p.id AND ul.user_id = s.user_id 
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE
				(
					p.author_id = s.user_id
					OR (p.privacy_mode = 1 AND u.private_mode = 0)
					OR (f.user_to IS NOT NULL AND p.privacy_mode IN (1,2))
					OR (p.privacy_mode = 3 AND pr.follower_id IS NOT NULL)
					OR (p.group_id IS NOT NULL AND gmr.group_id IS NOT NULL)
				)
				AND p.id < ?
			GROUP BY p.id, p.author_id, u.nick_name, u.avatar, p.message, p.image, p.date, p.privacy_mode, p.group_id, g.title, ul.id
			ORDER BY p.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"], obj["last_id"], obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PostFeed
	for rows.Next() {
		var postId, authorId, privacyMode, groupId, likeCount, commentCount, likeId int
		var nickname, message, date, groupTitle, postImage, avatar, lastName, firstName string
		var isClient bool

		err := rows.Scan(
			&postId,
			&authorId,
			&nickname,
			&firstName,
			&lastName,
			&avatar,
			&message,
			&postImage,
			&date,
			&privacyMode,
			&groupId,
			&likeCount,
			&commentCount,
			&groupTitle,
			&likeId,
			&isClient,
		)
		if err != nil {
			return nil, err
		}

		pf := models.PostFeed{
			Post: &models.Post{
				Id:          postId,
				Message:     message,
				Date:        date,
				PrivacyMode: privacyMode,
				Author: &models.User{
					Id:        authorId,
					FirstName: firstName,
					LastName:  lastName,
					IsClient:  isClient,
				},
				Group: &models.Group{
					Id:    groupId,
					Title: groupTitle,
				},
			},
			LikeCount:    likeCount,
			CommentCount: commentCount,
		}

		if likeId > 0 {
			pf.Like = &models.Like{
				Id: likeId,
			}
		}

		if nickname != "" {
			pf.Post.Author.Nickname = nickname
		}

		if postImage != "" {
			pf.Post.Image = &postImage
		}
		if avatar != "" {
			pf.Post.Author.Avatar = &avatar
		}

		result = append(result, pf)
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
				COALESCE(u.nick_name, ''),
				u.first_name,
				u.last_name,
				COALESCE(u.avatar, ''),
				p.message,
				COALESCE(p.image, ''),
				p.date,
				p.privacy_mode,
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id),
				COALESCE(ul.id, 0),
				CASE
					WHEN p.author_id = s.user_id THEN 1
					ELSE 0
				END AS is_client
			FROM posts p
			JOIN users u ON p.author_id = u.id
			JOIN sessions s ON s.uuid = ?
			LEFT JOIN follow_rel f ON f.user_from = s.user_id AND f.user_to = p.author_id
			LEFT JOIN privacy_post_rel pr ON pr.post_id = p.id AND pr.follower_id = s.user_id
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN likes ul ON ul.post_id = p.id AND ul.user_id = s.user_id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE
				(
					f.user_to IS NOT NULL AND p.privacy_mode IN (1,2)
					OR (p.privacy_mode = 3 AND pr.follower_id IS NOT NULL)
				)
				AND p.id < ?
			GROUP BY p.id, p.author_id, u.nick_name, u.avatar, p.message, p.image, p.date, p.privacy_mode, ul.id
			ORDER BY p.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"], obj["last_id"], obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PostFeed
	for rows.Next() {
		var postId, authorId, privacyMode, likeCount, commentCount, likeId int
		var nickname, message, date, avatar, postImage, lastName, firstName string
		var isClient bool

		err := rows.Scan(
			&postId,
			&authorId,
			&nickname,
			&firstName,
			&lastName,
			&avatar,
			&message,
			&postImage,
			&date,
			&privacyMode,
			&likeCount,
			&commentCount,
			&likeId,
			&isClient,
		)
		if err != nil {
			return nil, err
		}

		pf := models.PostFeed{
			Post: &models.Post{
				Id:          postId,
				Message:     message,
				Date:        date,
				PrivacyMode: privacyMode,
				Author: &models.User{
					Id:        authorId,
					FirstName: firstName,
					LastName:  lastName,
					IsClient:  isClient,
				},
			},

			LikeCount:    likeCount,
			CommentCount: commentCount,
		}

		if likeId > 0 {
			pf.Like = &models.Like{
				Id: likeId,
			}
		}
		if nickname != "" {
			pf.Post.Author.Nickname = nickname
		}

		if postImage != "" {
			pf.Post.Image = &postImage
		}
		if avatar != "" {
			pf.Post.Author.Avatar = &avatar
		}

		result = append(result, pf)

	}
	return &models.Response{Result: result}, nil
}

func (db *PostDB) GetGroupFeed(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
			{
				session_uuid : string,
				group_id : int,
				last_id : int,
				limit : int
			}
	*/

	stmt := `SELECT 
				p.id,
				p.author_id,
				COALESCE(u.nick_name, ''),
				u.first_name,
				u.last_name,
				COALESCE(u.avatar, ''),
				p.message,
				COALESCE(p.image, ''),
				p.date,
				COUNT(DISTINCT l.id),
				COUNT(DISTINCT c.id),
				COALESCE(ul.id, 0),
				CASE
					WHEN p.author_id = s.user_id THEN 1
					ELSE 0
				END AS is_client
			FROM posts p
			JOIN users u ON p.author_id = u.id
			LEFT JOIN sessions s ON s.uuid = ?
			LEFT JOIN likes l ON l.post_id = p.id
			LEFT JOIN likes ul ON ul.post_id = p.id AND ul.user_id = s.user_id
			LEFT JOIN comments c ON c.post_id = p.id
			WHERE	
				p.group_id = ?
				AND p.id < ?
			GROUP BY p.id, p.author_id, u.nick_name, u.avatar, p.message, p.image, p.date, ul.id
			ORDER BY p.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"], obj["group_id"], obj["last_id"], obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PostFeed
	for rows.Next() {
		var postId, authorId, likeCount, commentCount, likeId int
		var nickname, message, date, postImage, avatar, firstName, lastName string
		var isClient bool

		err := rows.Scan(
			&postId,
			&authorId,
			&nickname,
			&firstName,
			&lastName,
			&avatar,
			&message,
			&postImage,
			&date,
			&likeCount,
			&commentCount,
			&likeId,
			&isClient,
		)
		if err != nil {
			return nil, err
		}
		pf := models.PostFeed{
			Post: &models.Post{
				Id:      postId,
				Message: message,
				Date:    date,
				Author: &models.User{
					Id:        authorId,
					FirstName: firstName,
					LastName:  lastName,
					IsClient:  isClient,
				},
			},
			LikeCount:    likeCount,
			CommentCount: commentCount,
		}

		if likeId > 0 {
			pf.Like = &models.Like{
				Id: likeId,
			}
		}
		if nickname != "" {
			pf.Post.Author.Nickname = nickname
		}

		if postImage != "" {
			pf.Post.Image = &postImage
		}
		if avatar != "" {
			pf.Post.Author.Avatar = &avatar
		}

		result = append(result, pf)
	}
	return &models.Response{Result: result}, nil
}
