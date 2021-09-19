package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	"go.uber.org/zap"
)

// Create User
func (m *manager) Create(u *models.User) (*models.User, error) {
	if err := m.db.Create(u).Error; err != nil {
		m.log.Error(err.Error(), zap.Any("gorm", "create"))
		return nil, err
	}
	return u, nil
}
