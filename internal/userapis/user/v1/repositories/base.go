package repositories

import "github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"

func (m *manager) NewSort(key string, is_asc bool) models.Sort {
	return models.Sort{
		Key:   key,
		IsASC: is_asc,
	}
}

func (m *manager) NewFilter(key string, value string, method string) models.Filter {
	return models.Filter{
		Key:    key,
		Value:  value,
		Method: method,
	}
}

func (m *manager) NewPagination(limit uint32, page uint32) models.Pagination {
	return models.Pagination{
		Limit: limit,
		Page:  page,
	}
}
