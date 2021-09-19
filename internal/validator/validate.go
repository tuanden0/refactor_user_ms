package validator

import (
	"sync"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	validateOnce sync.Once
	Valid        Validator
)

type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
	log      zap.Logger
}

func NewValidator(log zap.Logger) Validator {
	validateOnce.Do(func() {
		v := validator.New()

		valid := Validator{
			Validate: v,
			log:      log,
		}
		Valid = valid
	})

	return Valid
}
