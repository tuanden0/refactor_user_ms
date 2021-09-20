package helpers

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func MapUpdateRequest(ctx context.Context, in *userV1PB.UpdateRequest) (*models.User, error) {

	u := &models.User{
		ID:       in.GetId(),
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
		Role:     uint32(in.GetRole()),
	}

	return u, nil
}

func MapUpdateResponse(u *models.User) *userV1PB.UpdateResponse {
	return &userV1PB.UpdateResponse{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
		Role:     userV1PB.Role(u.GetRole()),
	}
}
