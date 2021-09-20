package repositories

import (
	"fmt"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	"go.uber.org/zap"
)

func (m *manager) List(pg *models.Pagination, s *models.Sort, fs []*models.Filter) ([]*models.User, error) {

	users := make([]*models.User, 0)

	qs := m.db.
		Order(fmt.Sprintf("%v %v", s.GetKey(), s.GetIsASC())).
		Limit(int(pg.GetLimit())).
		Offset(int(pg.GetLimit()) * (int(pg.GetPage()) - 1))

	for _, f := range fs {
		qs = qs.Where(fmt.Sprintf("%v %v ?", f.GetKey(), f.GetMethod()), f.GetValue())
	}

	if err := qs.Find(&users).Error; err != nil {
		m.log.Error(err.Error(), zap.String("gorm", "list"))
		return nil, err
	}

	return users, nil
}
