package repositories

import (
	"gorm.io/gorm"
)

type manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) UserRepository {
	return &manager{
		db: db,
	}
}
