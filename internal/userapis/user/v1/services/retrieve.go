package services

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/helpers"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
	"go.uber.org/zap"
)

func (s *service) Retrieve(ctx context.Context, in *userV1PB.RetrieveRequest) (*userV1PB.RetrieveResponse, error) {

	// Validate RetrieveRequest
	if err := s.validator.RetrieveRequest(ctx, in); err != nil {
		s.log.Error("user input invalid", zap.Any("retrieve_user_validate_input", err))
		return nil, err
	}

	// Mapping RetrieveRequest
	id := helpers.MapRetrieveRequest(ctx, in)

	// Fetch user data
	u, err := s.repo.Retrieve(id)
	if err != nil {
		return nil, helpers.RetrieveError
	}

	return helpers.MapRetrieveResponse(u), nil
}
