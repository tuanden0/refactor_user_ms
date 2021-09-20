package services

import (
	"context"

	logger "github.com/tuanden0/refactor_user_ms/internal/logs/zap_driver"
	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/helpers"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

func (s *service) Delete(ctx context.Context, in *userV1PB.DeleteRequest) (*userV1PB.DeleteResponse, error) {

	// Validate DeleteRequest
	if err := s.validator.DeleteRequest(ctx, in); err != nil {
		logger.Error("user input invalid", zap.String("delete_user_validate_input", err.Error()))
		return nil, err
	}

	// Mapping DeleteRequest
	id := helpers.MapDeleteRequest(ctx, in)

	// Delete user in database
	if err := s.repo.Delete(id); err != nil {
		logger.Error("failed to delete user", zap.String("delete_user_error", err.Error()))
		return nil, helpers.DeleteError
	}

	return helpers.MapDeleteResponse(), nil
}
