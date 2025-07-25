package models_group

import "social-network/pkg/db/models"

type GroupDB struct {
	*models.DB
}

func New(db *models.DB) *GroupDB {
	return &GroupDB{DB: db}
}
