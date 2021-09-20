package services

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/helpers"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

func (s *service) List(ctx context.Context, in *userV1PB.ListRequest) (*userV1PB.ListResponse, error) {

	// Validate ListRequest
	if err := s.validator.ListRequest(ctx, in); err != nil {
		s.log.Error("user input invalid", zap.Any("list_user_validate_input", err))
	}

	// Mapping ListRequest
	pg, sort, fs := helpers.MapListRequest(ctx, in)

	// Fetch users
	us, err := s.repo.List(pg, sort, fs)
	if err != nil {
		return nil, helpers.ListError
	}

	return helpers.MapListResponse(us), nil
}
