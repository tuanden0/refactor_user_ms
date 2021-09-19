package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
