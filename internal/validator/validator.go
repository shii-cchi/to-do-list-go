package validator

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func InitValidator() (*validator.Validate, error) {
	v := validator.New()
	if err := v.RegisterValidation("rfc3339", validateRFC3339); err != nil {
		return nil, err
	}

	return v, nil
}

func validateRFC3339(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
