package validation

import (
	"encoding/json"
	"errors"

	"github.com/diogoalvesf/my-routine-backend/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")

		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid JSON format: " + jsonErr.Error())
	} else if errors.As(validation_err, &jsonValidationError) {
		errorMessages := []rest_err.Causes{}

		for _, fieldErr := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Field:   fieldErr.Field(),
				Message: fieldErr.Translate(transl),
			}

			errorMessages = append(errorMessages, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorMessages)
	} else {
		return rest_err.NewBadRequestError("Invalid request data: " + validation_err.Error())
	}
}
