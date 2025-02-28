package validation

import (
	"encoding/json"
	"errors"

	"github.com/HunCoding/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

func initializeValidator() (*validator.Validate, ut.Translator, error) {
	validate := validator.New()
	enLocale := en.New()
	uniTrans := ut.New(enLocale, enLocale)

	translator, ok := uniTrans.GetTranslator("en")
	if !ok {
		return nil, nil, errors.New("unable to get translator for 'en' locale")
	}

	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en_translation.RegisterDefaultTranslations(val, translator)
	}

	return validate, translator, nil
}

func ValidateUserError(validationErr error, translator ut.Translator) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Tipo de campo inválido")
	}

	if errors.As(validationErr, &validationErrors) {
		return formatValidationErrors(validationErrors, translator)
	}

	return rest_err.NewBadRequestError("Erro ao tentar converter os campos")
}

func formatValidationErrors(validationErrors validator.ValidationErrors, translator ut.Translator) *rest_err.RestErr {
	errorCauses := make([]rest_err.Cause, len(validationErrors))

	for i, e := range validationErrors {
		cause := rest_err.Cause{
			Message: e.Translate(translator),
			Field:   e.Field(),
		}
		errorCauses[i] = cause
	}

	return rest_err.NewBadRequestValidationError("Alguns campos são inválidos", errorCauses)
}

func ValidarUsandoValidador(obj interface{}) *rest_err.RestErr {
	validate, translator, err := initializeValidator()
	if err != nil {
		// return rest_err.NewInternalServerError("Erro interno ao inicializar validador")
	}

	err = validate.Struct(obj)
	if err != nil {
		return ValidateUserError(err, translator)
	}

	return nil
}
