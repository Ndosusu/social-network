package models_like

import "social-network/pkg/db/models"

type LikeDB struct {
	*models.DB
}

func New(db *models.DB) *LikeDB {
	return &LikeDB{DB: db}
}
