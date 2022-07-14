package structs

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

//ErrorResponse .
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

//ValidateStruct .
func ValidateStruct(s interface{}) []*ErrorResponse {
	var errorsResponse []*ErrorResponse
	err := validate.Struct(s)
	var validationErrors validator.ValidationErrors
	if err != nil && errors.As(err, &validationErrors) {
		for _, err := range validationErrors {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errorsResponse = append(errorsResponse, &element)
		}
	}
	return errorsResponse
}
