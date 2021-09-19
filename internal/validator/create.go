package validator

import (
	"context"

	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func (v *Validator) CreateRequest(ctx context.Context, in *userV1PB.CreateRequest) error {

	if err := v.Validate.Struct(in); err != nil {
		return v.ParseError(err)
	}

	return nil

}
