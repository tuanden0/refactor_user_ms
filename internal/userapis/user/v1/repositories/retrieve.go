package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
)

func (m *manager) Retrieve(id uint64) (*models.User, error) {

	user := models.User{}

	if err := m.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
