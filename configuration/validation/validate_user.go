package validation

import (
	"encoding/json"
	"errors"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate  = validator.New()
	translate ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translate, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, translate)
	}
}

func ValidateUserError(
	validation_err error,
) *error_handle.ErrorHandle {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return error_handle.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []error_handle.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := error_handle.Causes{
				Message: e.Translate(translate),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return error_handle.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return error_handle.NewBadRequestError("Error trying to convert fields")
	}
}
