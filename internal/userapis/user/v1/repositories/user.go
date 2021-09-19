package repositories

import "github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"

type UserRepository interface {
	Create(u *models.User) (*models.User, error)
	Retrieve(id string) (*models.User, error)
	Update(id string, in models.User) (*models.User, error)
	Delete(id string) error
	List(pagination models.Pagination, sort models.Sort, filters []models.Filter) ([]*models.User, error)
}
