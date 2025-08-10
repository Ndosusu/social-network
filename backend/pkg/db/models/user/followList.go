package models_user

import (
	"fmt"
	"social-network/pkg/db/models"
)

func (db *UserDB) GetFollowFrom(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			user_id : int,
			last_id : int,
			limit : int,
			client_id : int,
		}
	*/
	stmt := `SELECT
				u.id,
				u.first_name,
				u.last_name,
				COALESCE(u.nick_name, ''),
				COALESCE(u.avatar, ''),
				u.private_mode,
				CASE
					WHEN u.id = ?
					THEN 1
					ELSE 0
				END AS is_client
			FROM follow_rel fr
			JOIN users u ON u.id = fr.user_to 
			WHERE 
				fr.user_from = ?
				AND fr.user_to > ?
			ORDER BY u.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["client_id"], obj["user_id"], obj["last_id"], obj["limit"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []models.User
	for rows.Next() {
		var userID int
		var firstName, lastName, nickName, avatar string
		var privateMode, isClient bool

		err := rows.Scan(&userID, &firstName, &lastName, &nickName, &avatar, &privateMode, &isClient)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		user := models.User{
			Id:          userID,
			FirstName:   firstName,
			LastName:    lastName,
			PrivateMode: privateMode,
			IsClient:    isClient,
		}
		if nickName != "" {
			user.Nickname = nickName
		}
		if avatar != "" {
			user.Avatar = &avatar
		}

		result = append(result, user)
	}
	return &models.Response{Result: result}, nil
}

func (db *UserDB) GetFollowTo(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			user_id : int,
			last_id : int,
			limit : int,
			client_id : int,
		}
	*/
	stmt := `SELECT
				u.id,
				u.first_name,
				u.last_name,
				COALESCE(u.nick_name, ''),
				COALESCE(u.avatar, ''),
				u.private_mode,
				CASE
					WHEN u.id = ?
					THEN 1
					ELSE 0
				END AS is_client
			FROM follow_rel fr
			JOIN users u ON u.id = fr.user_from
			WHERE 
				fr.user_to = ?
				AND fr.user_from > ?
			ORDER BY u.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["client_id"], obj["user_id"], obj["last_id"], obj["limit"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []models.User
	for rows.Next() {
		var userID int
		var firstName, lastName, nickName, avatar string
		var privateMode, isClient bool

		err := rows.Scan(&userID, &firstName, &lastName, &nickName, &avatar, &privateMode, &isClient)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		user := models.User{
			Id:          userID,
			FirstName:   firstName,
			LastName:    lastName,
			PrivateMode: privateMode,
			IsClient:    isClient,
		}
		if nickName != "" {
			user.Nickname = nickName
		}
		if avatar != "" {
			user.Avatar = &avatar
		}

		result = append(result, user)
	}
	return &models.Response{Result: result}, nil
}
