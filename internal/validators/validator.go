package validators

import (
	"sync"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (v *Validator) InitValidate() {}

func (v *Validator) ParseError(err error) error {

	errs := err.(validator.ValidationErrors)
	st := status.New(codes.InvalidArgument, "invalid_input")
	br := &errdetails.BadRequest{}

	for _, e := range errs {
		v := &errdetails.BadRequest_FieldViolation{
			Field:       e.Field(),
			Description: e.Translate(v.Trans),
		}
		br.FieldViolations = append(br.FieldViolations, v)
	}

	st, err = st.WithDetails(br)
	if err != nil {
		v.log.Error("Unexpected error attaching metadata", zap.Any("validator_parse_error", err.Error()))
	}

	return st.Err()
}
