package models_group

import (
	"social-network/pkg/db/models"
)

func (db *GroupDB) GetUserGroupList(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			session_uuid : string,
		}
	*/
	stmt := `SELECT
				g.id, 
				g.title, 
				CASE
					WHEN g.admin_id = s.user_id THEN 1
					ELSE 0 
				END as is_admin,
			FROM groups g
			JOIN sessions s on s.uuid = ?
			LEFT JOIN group_member_rel gmr ON gm.group_id = g.id 
			WHERE
				gmr.member_id = s.user_id
			GROUP BY g.id, g.title, is_admin
			ORDER BY g.id DESC;`
	rows, err := db.Conn.Query(stmt, obj["session_uuid"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Group

	for rows.Next() {
		var groupID int
		var groupTitle string
		var isAdmin bool

		err := rows.Scan(
			&groupID,
			&groupTitle,
			&isAdmin,
		)
		if err != nil {
			return nil, err
		}

		group := models.Group{
			Id:    groupID,
			Title: groupTitle,
			Admin: &models.User{
				IsClient: isAdmin,
			},
		}

		result = append(result, group)

	}
	return &models.Response{Result: result}, nil
}
