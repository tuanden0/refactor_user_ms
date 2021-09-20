package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
)

// Create User
func (m *manager) Create(u *models.User) (*models.User, error) {
	if err := m.db.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
