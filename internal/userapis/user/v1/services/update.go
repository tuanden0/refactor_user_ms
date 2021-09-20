package services

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/helpers"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

func (s *service) Update(ctx context.Context, in *userV1PB.UpdateRequest) (*userV1PB.UpdateResponse, error) {

	// Validate UpdateRequest
	if err := s.validator.UpdateRequest(ctx, in); err != nil {
		s.log.Error("user input invalid", zap.Any("update_user_validate_input", err))
		return nil, err
	}

	// Mapping UpdateRequest
	u, err := helpers.MapUpdateRequest(ctx, in)
	if err != nil {
		s.log.Error("failed to map user input", zap.String("update_user_map_error", err.Error()))
		return nil, helpers.MappingError
	}

	// Update data to db
	user, err := s.repo.Update(u.GetID(), u)
	if err != nil {
		return nil, helpers.UpdateError
	}

	return helpers.MapUpdateResponse(user), nil
}
