package validators

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func (v *Validator) InitTranslator() error {

	en := en.New()
	uni := ut.New(en, en)

	trans, found := uni.GetTranslator("en")
	if !found {
		return fmt.Errorf("translator not found")
	}

	v.Trans = trans

	if err := en_translations.RegisterDefaultTranslations(v.Validate, trans); err != nil {
		return err
	}

	// Get lower-case field name
	v.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register custom error msg translator
	v.registerCustomTranslate(trans, "required", "{0} is a required field")
	v.registerCustomTranslate(trans, "email", "{0} must be a valid email")

	return nil
}

func (v *Validator) registerCustomTranslate(trans ut.Translator, tag, msg string) {

	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, msg, true)
	}

	translationFn := func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.StructField())
		return t
	}

	v.Validate.RegisterTranslation(tag, trans, registerFn, translationFn)
}
