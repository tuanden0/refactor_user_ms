package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
)

func (m *manager) Delete(id uint64) error {

	if err := m.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
