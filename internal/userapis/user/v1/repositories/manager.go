package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type manager struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewManager(db *gorm.DB, log *zap.Logger) UserRepository {
	return &manager{
		db:  db,
		log: log,
	}
}
