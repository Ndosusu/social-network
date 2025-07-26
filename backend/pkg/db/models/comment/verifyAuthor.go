package models_comment

import "social-network/pkg/db/models"

func (db *CommentDB) IsUserCommentAuthor(obj map[string]any) (*models.Response, error) {
	/*
	   expected input:
	   {
	       comment_id: int,
	       session_uuid: string,
	   }
	*/
	stmt := `SELECT EXISTS(
        SELECT 1 FROM comments c
        JOIN sessions s ON s.uuid = ?
        WHERE c.id = ? 
        AND c.author_id = s.user_id
        AND (s.date_expiration IS NULL OR s.date_expiration >= CURRENT_TIMESTAMP)
    );`

	var canDelete bool
	err := db.Conn.QueryRow(stmt, obj["session_uuid"], obj["comment_id"]).Scan(&canDelete)
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: canDelete}, nil
}
