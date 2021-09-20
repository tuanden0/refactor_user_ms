package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	"go.uber.org/zap"
)

func (m *manager) Retrieve(id uint64) (*models.User, error) {

	user := models.User{}

	if err := m.db.First(&user, id).Error; err != nil {
		m.log.Error(err.Error(), zap.String("gorm", "retrieve"))
		return nil, err
	}

	return &user, nil
}
