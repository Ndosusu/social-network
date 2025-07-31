package models_user

import "social-network/pkg/db/models"

func (db *UserDB) SearchUser(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			query : string,
			limit : int
		}
	*/
	searchPattern := "%" + obj["query"].(string) + "%"
	stmt := `SELECT
				u.id,
				COALESCE(u.nick_name, ''),
				u.first_name,
				u.last_name,
				COALESCE(u.avatar, ''),
				u.private_mode,
			FROM users u
			LEFT JOIN sessions s ON s.user_id = u.id
			WHERE
				(
					LOWER(u.nick_name) LIKE LOWER(?)
					OR LOWER(u.first_name) LIKE LOWER(?)
					OR LOWER(u.last_name) LIKE LOWER(?)
					OR LOWER(u.about) LIKE LOWER(?)
				)
			GROUP BY u.id, u.nick_name, u.first_name, u.last_name, u.avatar
			ORDER BY u.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, searchPattern, searchPattern, searchPattern, obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var userId int
		var nickname, firstName, lastName, avatar string
		var privateMode bool

		err := rows.Scan(&userId, &nickname, &firstName, &lastName, &avatar, &privateMode)
		if err != nil {
			continue
		}

		user := models.User{
			Id:          userId,
			FirstName:   firstName,
			LastName:    lastName,
			PrivateMode: privateMode,
		}

		if avatar != "" {
			user.Avatar = &avatar
		}
		if nickname != "" {
			user.Nickname = nickname
		}

		users = append(users, user)
	}

	return &models.Response{Result: users}, nil

}
