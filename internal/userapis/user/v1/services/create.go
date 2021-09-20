package services

import (
	"context"

	logger "github.com/tuanden0/refactor_user_ms/internal/logs/zap_driver"
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/helpers"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

func (s *service) Create(ctx context.Context, in *userV1PB.CreateRequest) (*userV1PB.CreateResponse, error) {

	// Validate CreateRequest
	if err := s.validator.CreateRequest(ctx, in); err != nil {
		logger.Error("user input invalid", zap.String("create_user_validate_input", err.Error()))
		return nil, err
	}

	// Mapping data to User struct
	u, mapErr := helpers.MapCreateRequest(ctx, in)
	if mapErr != nil {
		logger.Error("failed to map user input", zap.String("create_user_map_error", mapErr.Error()))
		return nil, helpers.MappingError
	}

	// Add data to database
	user, createErr := s.repo.Create(u)
	if createErr != nil {
		logger.Error("failed to create user", zap.String("create_user_error", createErr.Error()))
		return nil, helpers.CreateError
	}

	return helpers.MapCreateResponse(user), nil
}
