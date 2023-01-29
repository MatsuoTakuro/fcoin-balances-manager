package api

import (
	"github.com/go-playground/validator/v10"
)

func invalidParams(v *validator.Validate, err error) []string {
	var invalidParams []string
	for _, err := range err.(validator.ValidationErrors) {
		p := err.Field()
		invalidParams = append(invalidParams, p)
	}
	return invalidParams
}
