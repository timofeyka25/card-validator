package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation error: field '%s' failed on the '%s' tag", err.Field(), err.Tag())
		}
	}
	return nil
}
