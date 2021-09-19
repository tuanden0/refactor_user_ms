package helpers

import (
	"context"

	"github.com/tuanden0/refactor_user_ms/internal/userapis/user/v1/models"
	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func MapCreateRequest(ctx context.Context, in *userV1PB.CreateRequest) (*models.User, error) {

	u := &models.User{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
		Role:     in.GetRole().String(),
	}

	hash, err := u.HashPassword()
	if err != nil {
		return nil, err
	}
	u.Password = hash

	return u, nil
}

func MapCreateResponse(u *models.User) *userV1PB.CreateResponse {
	return &userV1PB.CreateResponse{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
		Role:     userV1PB.Role(userV1PB.Role_value[u.GetRole()]),
	}
}
