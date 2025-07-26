package models_comment

import "social-network/pkg/db/models"

type CommentDB struct {
	*models.DB
}

func New(db *models.DB) *CommentDB {
	return &CommentDB{DB: db}
}
