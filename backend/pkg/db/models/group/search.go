package models_group

import "social-network/pkg/db/models"

func (db *GroupDB) SearchGroup(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			query : string,
			limit : int
		}
	*/
	searchPattern := "%" + obj["query"].(string) + "%"
	stmt := `SELECT
				g.id,
				g.title,
			FROM groups g
			WHERE
				(
					LOWER(g.title) LIKE LOWER(?)
					OR LOWER(g.about) LIKE LOWER(?)
				)
			GROUP BY u.id, u.nick_name, u.first_name, u.last_name, u.avatar
			ORDER BY u.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, searchPattern, searchPattern, obj["limit"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var groupID int
		var groupTitle string

		err := rows.Scan(&groupID, &groupTitle)
		if err != nil {
			continue
		}

		user := models.Group{
			Id:    groupID,
			Title: groupTitle,
		}

		groups = append(groups, user)
	}

	return &models.Response{Result: groups}, nil

}
