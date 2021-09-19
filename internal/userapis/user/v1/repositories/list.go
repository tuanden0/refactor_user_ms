package repositories

import (
	"fmt"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	"go.uber.org/zap"
)

func (m *manager) List(pagination models.Pagination, sort models.Sort, filters []models.Filter) ([]*models.User, error) {

	users := make([]*models.User, 0)

	qs := m.db.
		Order(fmt.Sprintf("%v %v", sort.GetKey(), sort.GetIsASC())).
		Limit(int(pagination.GetLimit())).
		Offset(int(pagination.GetLimit()) * (int(pagination.GetPage()) - 1))

	for _, f := range filters {
		qs = qs.Where(fmt.Sprintf("%v %v ?", f.GetKey(), f.GetMethod()), f.GetValue())
	}

	if err := qs.Find(&users).Error; err != nil {
		m.log.Error(err.Error(), zap.Any("gorm", "list"))
		return nil, err
	}

	return users, nil
}
