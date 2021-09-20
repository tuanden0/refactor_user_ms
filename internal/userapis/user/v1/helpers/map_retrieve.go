package helpers

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func MapRetrieveRequest(ctx context.Context, in *userV1PB.RetrieveRequest) uint64 {
	return in.GetId()
}

func MapRetrieveResponse(u *models.User) *userV1PB.RetrieveResponse {
	return &userV1PB.RetrieveResponse{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
		Role:     userV1PB.Role(u.GetRole()),
	}
}
