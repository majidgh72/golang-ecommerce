package validate

import (
	"github.com/go-playground/validator/v10"
)

func Check(data interface{}) (bool, map[string]string) {
	validate := validator.New()

	errs := validate.Struct(data)

	passed := true
	errMsgs := make(map[string]string)

	if errs != nil {
		passed = false
		for _, err := range errs.(validator.ValidationErrors) {
			errMsgs[err.Field()] = "Error placed here"
		}
	}

	return passed, errMsgs
}
