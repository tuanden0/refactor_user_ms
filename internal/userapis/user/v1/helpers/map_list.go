package helpers

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func MapListRequest(ctx context.Context, in *userV1PB.ListRequest) (*models.Pagination, *models.Sort, []*models.Filter) {

	// Get input
	s := in.GetSort()
	fs := in.GetFilters()
	pg := in.GetPagination()

	// Mapping input to models
	sort := &models.Sort{
		Key:   s.GetKey(),
		IsASC: s.GetIsAsc(),
	}

	pagination := &models.Pagination{
		Limit: uint32(pg.GetLimit()),
		Page:  uint32(pg.GetPage()),
	}

	filters := make([]*models.Filter, 0)
	for _, f := range fs {
		filters = append(filters, &models.Filter{
			Key:    f.GetKey(),
			Value:  f.GetValue(),
			Method: f.GetMethod(),
		})
	}

	return pagination, sort, filters
}

func MapListResponse(us []*models.User) *userV1PB.ListResponse {

	res := make([]*userV1PB.UserList, 0)
	for _, u := range us {
		res = append(res, &userV1PB.UserList{
			Id:       u.GetID(),
			Username: u.GetUserName(),
			Email:    u.GetEmail(),
			Role:     userV1PB.Role(u.GetRole()),
		})
	}

	return &userV1PB.ListResponse{
		Users: res,
	}
}
