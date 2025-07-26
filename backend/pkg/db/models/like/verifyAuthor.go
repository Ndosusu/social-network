package models_like

import "social-network/pkg/db/models"

func (db *LikeDB) IsUserLikeAuthor(obj map[string]any) (*models.Response, error) {
	/*
	   expected input:
	   {
	       like_id: int,
	       session_uuid: string,
	   }
	*/
	stmt := `SELECT EXISTS(
        SELECT 1 FROM likes l
        JOIN sessions s ON s.uuid = ?
        WHERE l.id = ? 
        AND l.user_id = s.user_id
        AND (s.date_expiration IS NULL OR s.date_expiration >= CURRENT_TIMESTAMP)
    );`

	var canDelete bool
	err := db.Conn.QueryRow(stmt, obj["session_uuid"], obj["like_id"]).Scan(&canDelete)
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: canDelete}, nil
}
