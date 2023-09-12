package middlewares

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type RequestBodyValidationError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

var validate = validator.New()

func validateRequestBody(data interface{}) []RequestBodyValidationError {
	var validationErrors []RequestBodyValidationError

	errs := validate.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem RequestBodyValidationError

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func RequestBodyValidatorMiddleware(validationDTO interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {

		jsonParsingError := c.BodyParser(validationDTO)

		if jsonParsingError != nil {
			// handle error
			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: jsonParsingError.Error(),
			}
		}

		errors := validateRequestBody(validationDTO)

		ErrorIsPresent := errors != nil && len(errors) > 0 && errors[0].Error

		if ErrorIsPresent {

			errorMessages := make([]string, 0)

			for _, err := range errors {
				errorMessages = append(errorMessages, fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					err.FailedField,
					err.Value,
					err.Tag,
				))
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errorMessages, " and "),
			}
		}

		return c.Next()

	}
}
