package models_post

import "social-network/pkg/db/models"

func (db *PostDB) IsUserPostAuthor(obj map[string]any) (*models.Response, error) {
	/*
	   expected input:
	   {
	       post_id: int,
	       session_uuid: string,
	   }
	*/
	stmt := `SELECT EXISTS(
        SELECT 1 FROM posts p
        JOIN sessions s ON s.uuid = ?
        WHERE p.id = ? 
        AND p.author_id = s.user_id
        AND (s.date_expiration IS NULL OR s.date_expiration >= CURRENT_TIMESTAMP)
    );`

	var canDelete bool
	err := db.Conn.QueryRow(stmt, obj["session_uuid"], obj["post_id"]).Scan(&canDelete)
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: canDelete}, nil
}
