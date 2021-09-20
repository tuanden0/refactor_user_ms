package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
)

func (m *manager) Update(id uint64, in *models.User) (*models.User, error) {

	user := models.User{}

	if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if err := m.db.Model(&user).Omit("id").Updates(&in).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
