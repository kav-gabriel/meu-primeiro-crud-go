package validation

import (
	"encoding/json"
	"errors"

	"github.com/HumCoding/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate *validator.Validate
	transl   ut.Translator
)

func init() {
	Validate = validator.New()
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		uniTrans := ut.New(enLocale, enLocale)
		transl, _ = uniTrans.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

// ValidateUserError translates and formats validation errors into a RestErr.
func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	// Check for JSON unmarshaling error
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	}

	// Check for validation errors
	if errors.As(validationErr, &validationErrors) {
		errorCauses := []rest_err.Cause{}

		for _, e := range validationErrors {
			cause := rest_err.Cause{
				Message: e.Translate(transl), // Translate error message
				Field:   e.Field(),           // Get the field name
			}
			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	}

	// Generic error for unexpected issues
	return rest_err.NewBadRequestError("Error trying to convert fields")
}
