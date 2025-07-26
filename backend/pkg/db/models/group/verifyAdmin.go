package models_group

import "social-network/pkg/db/models"

func (db *GroupDB) IsUserGroupAdmin(obj map[string]any) (*models.Response, error) {
	/*
	   expected input:
	   {
	       group_id: int,
	       session_uuid: string,
	   }
	*/
	stmt := `SELECT EXISTS(
        SELECT 1 FROM groups g
        JOIN sessions s ON s.uuid = ?
        WHERE g.id = ? 
        AND g.admin = s.user_id
        AND (s.date_expiration IS NULL OR s.date_expiration >= CURRENT_TIMESTAMP)
    );`

	var canDelete bool
	err := db.Conn.QueryRow(stmt, obj["session_uuid"], obj["group_id"]).Scan(&canDelete)
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: canDelete}, nil
}
