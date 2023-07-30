package validation

import (
	"encoding/json"
	"errors"

	"github.com/MogLuiz/go-person-api/src/configuration/error_logger"
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
) *error_logger.ErrorLogger {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return error_logger.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []error_logger.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := error_logger.Causes{
				Message: e.Translate(translate),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return error_logger.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return error_logger.NewBadRequestError("Error trying to convert fields")
	}
}
