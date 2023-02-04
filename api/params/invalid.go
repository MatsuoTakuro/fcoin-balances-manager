package params

import (
	"github.com/go-playground/validator/v10"
)

func InvalidBodyItems(v *validator.Validate, err error) []string {
	var items []string
	for _, err := range err.(validator.ValidationErrors) {
		i := err.Field()
		items = append(items, i)
	}
	return items
}
