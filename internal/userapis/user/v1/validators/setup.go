package validators

import (
	vd "github.com/tuanden0/refactor_user_ms/internal/validators"
)

type UserValidator struct {
	userValid vd.Validator
}

func NewUserValidator(v vd.Validator) *UserValidator {
	return &UserValidator{v}
}
