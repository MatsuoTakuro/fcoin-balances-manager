package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func InvalidItemsErrMessages(v *validator.Validate, err error) []string {
	var errMsgs []string
	for _, err := range err.(validator.ValidationErrors) {
		msg := fmt.Sprintf("%s (input:%v) violates %s(=%s) validation", err.Field(), err.Value(), err.Tag(), err.Param())
		errMsgs = append(errMsgs, msg)
	}
	return errMsgs
}
