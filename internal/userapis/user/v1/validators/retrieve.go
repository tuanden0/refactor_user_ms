package validators

import (
	"context"

	userV1PB "github.com/tuanden0/refactor_user_ms/proto/gen/go/user/v1"
)

func (v *UserValidator) RetrieveRequest(ctx context.Context, in *userV1PB.RetrieveRequest) error {

	if err := v.userValid.Validate.Struct(in); err != nil {
		return v.userValid.ParseError(err)
	}

	return nil

}
