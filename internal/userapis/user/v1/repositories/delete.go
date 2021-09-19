package repositories

import (
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	"go.uber.org/zap"
)

func (m *manager) Delete(id string) error {

	if err := m.db.Delete(&models.User{}, id).Error; err != nil {
		m.log.Error(err.Error(), zap.Any("gorm", "delete"))
		return err
	}

	return nil
}
