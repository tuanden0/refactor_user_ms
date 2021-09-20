package services

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/helpers"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

func (s *service) Delete(ctx context.Context, in *userV1PB.DeleteRequest) (*userV1PB.DeleteResponse, error) {

	// Validate DeleteRequest
	if err := s.validator.DeleteRequest(ctx, in); err != nil {
		s.log.Error("user input invalid", zap.Any("delete_user_validate_input", err))
		return nil, err
	}

	// Mapping DeleteRequest
	id := helpers.MapDeleteRequest(ctx, in)

	// Delete user in database
	if err := s.repo.Delete(id); err != nil {
		return nil, helpers.DeleteError
	}

	return helpers.MapDeleteResponse(), nil
}
