package helpers

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func InitValidator() *validator.Validate {
	return validator.New()
}
func TranslateError(err error, validate *validator.Validate) error {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	var errs []error
	if err != nil {
		validatorErrs := err.(validator.ValidationErrors)
		for _, e := range validatorErrs {
			translatedErr := fmt.Errorf(e.Translate(trans))
			errs = append(errs, translatedErr)
		}
		return errs[0]
	}
	return nil
}
