package repositories

import "github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"

type UserRepository interface {
	Create(u *models.User) (*models.User, error)
	Retrieve(id uint64) (*models.User, error)
	Update(id uint64, in *models.User) (*models.User, error)
	Delete(id uint64) error
	List(pagination *models.Pagination, sort *models.Sort, filters []*models.Filter) ([]*models.User, error)
}
