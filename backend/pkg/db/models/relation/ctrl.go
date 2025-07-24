package models_rel

import "social-network/pkg/db/models"

type RelDB struct {
	*models.DB
}

func New(db *models.DB) *RelDB {
	return &RelDB{DB: db}
}
