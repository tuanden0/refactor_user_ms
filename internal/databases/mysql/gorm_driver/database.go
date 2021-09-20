package gormdriver

import (
	"fmt"
	"sync"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB         *gorm.DB
	DBErr      error
	dbConnOnce sync.Once
)

func ConnectDatabase() (*gorm.DB, error) {

	dbConnOnce.Do(func() {
		db, err := gorm.Open(sqlite.Open("local/gorm/database.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			DBErr = fmt.Errorf("failed to connect to database")
		}

		// Auto Migrate
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			DBErr = fmt.Errorf("failed to migrate database")
		}

		DB = db
	})

	if DBErr != nil {
		return nil, DBErr
	}

	return DB, nil
}
