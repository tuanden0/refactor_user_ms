package helpers

import (
	"context"

	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func MapDeleteRequest(ctx context.Context, in *userV1PB.DeleteRequest) uint64 {
	return in.GetId()
}

func MapDeleteResponse() *userV1PB.DeleteResponse {
	return &userV1PB.DeleteResponse{}
}
