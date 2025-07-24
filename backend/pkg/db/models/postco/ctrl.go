package models_post

import "social-network/pkg/db/models"

type PostDB struct {
	*models.DB
}

func New(db *models.DB) *PostDB {
	return &PostDB{DB: db}
}
